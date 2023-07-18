package processor

import (
	"ssv/internal/core"
	"ssv/internal/queue"
	"sync"
	"time"
)

type processor struct {
	queue Queue
	// chans map[int]map[int]chan queue.Task
	chans        map[int]map[string]chan string
	work         Work
	m            sync.Mutex
	workDuration time.Duration
}

func NewProcessor(queue Queue, work Work) *processor {
	return &processor{
		chans:        make(map[int]map[string]chan string, 0),
		queue:        queue,
		work:         work,
		workDuration: 3 * time.Second,
	}

}

// setWorkDuration sets the duration of the work
// Use for test
func (p *processor) setWorkDuration(duration time.Duration) {
	p.workDuration = duration
}

func (p *processor) run(ch chan string, validator int, dutyType string) {
	var task *queue.Task
	go func() {
		for message := range ch {
			p.work.Do(message)
			time.Sleep(p.workDuration)
			task = p.queue.GetTask(validator, dutyType)
			for task != nil {
				p.work.Do(task.Message)
				time.Sleep(p.workDuration)
				task = p.queue.GetTask(validator, dutyType)
			}
		}
	}()
}

func (p *processor) AddTask(validator int, dutyType string, task queue.Task) {
	p.m.Lock()
	defer p.m.Unlock()
	if _, ok := p.chans[validator]; !ok {
		p.chans[validator] = make(map[string]chan string)
		for _, duty := range core.DutyList {
			p.chans[validator][duty] = make(chan string)
			p.run(p.chans[validator][duty], validator, duty)
		}
		p.chans[validator][dutyType] <- task.Message
	} else {
		select {
		case p.chans[validator][dutyType] <- task.Message:
		default:
			p.queue.AddTask(validator, dutyType, task)
		}
	}
}

func (p *processor) Close() {
	for _, validator := range p.chans {
		for _, dutyType := range validator {
			close(dutyType)
		}
	}
}
