package queue

import (
	"container/heap"
	"ssv/internal/core"
	"sync"
)

type Task struct {
	Height  int
	Message string
}

type taskQueue struct {
	queue map[int]map[string]TaskHeap
	mutex sync.Mutex
}

func NewQueue() *taskQueue {
	return &taskQueue{}
}

func (q *taskQueue) AddTask(validator int, taskType string, task Task) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if _, ok := q.queue[validator]; !ok {
		q.queue = make(map[int]map[string]TaskHeap)
		q.queue[validator] = make(map[string]TaskHeap, len(core.DutyList))
		for _, dutyType := range core.DutyList {
			h := make(TaskHeap, 0)
			heap.Init(&h)
			q.queue[validator][dutyType] = h
		}
	}

	pQueue := q.queue[validator][taskType]
	heap.Push(&pQueue, task)
	q.queue[validator][taskType] = pQueue
}

func (q *taskQueue) GetTask(validator int, taskType string) *Task {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if _, ok := q.queue[validator]; !ok {
		return nil
	}
	if len(q.queue[validator][taskType]) == 0 {
		return nil
	}
	queue := q.queue[validator][taskType]
	task := queue.Pop()
	q.queue[validator][taskType] = queue
	t := task.(*Task)
	return t
}
