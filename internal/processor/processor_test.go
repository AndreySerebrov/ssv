package processor

import (
	"ssv/internal/core"
	"ssv/internal/queue"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
)

func TestSimpleOneTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	q := NewMockQueue(ctrl)
	w := NewMockWork(ctrl)
	p := NewProcessor(q, w)
	p.setWorkDuration(time.Millisecond * 300)
	defer p.Close()
	defer ctrl.Finish()
	//Expectations
	w.EXPECT().Do("a").Times(1)
	q.EXPECT().GetTask(0, gomock.Any()).Return(nil).MaxTimes(len(core.DutyList))

	p.AddTask(0, core.DutyList[0], queue.Task{Height: 1, Message: "a"})
	time.Sleep(time.Millisecond * 400)
}

func TestThreeTasksForOneValidator(t *testing.T) {
	ctrl := gomock.NewController(t)
	q := NewMockQueue(ctrl)
	w := NewMockWork(ctrl)
	p := NewProcessor(q, w)
	p.setWorkDuration(time.Millisecond * 100)
	defer p.Close()
	defer ctrl.Finish()

	//Expectations
	w.EXPECT().Do("a").Times(1)
	w.EXPECT().Do("b").Times(1)
	w.EXPECT().Do("c").Times(1)
	q.EXPECT().AddTask(0, core.DutyList[0], queue.Task{Height: 2, Message: "b"}).Times(1)
	q.EXPECT().AddTask(0, core.DutyList[0], queue.Task{Height: 3, Message: "c"}).Times(1)
	q.EXPECT().GetTask(0, core.DutyList[0]).Return(&queue.Task{Height: 2, Message: "b"}).Times(1)
	q.EXPECT().GetTask(0, core.DutyList[0]).Return(&queue.Task{Height: 3, Message: "c"}).Times(1)
	q.EXPECT().GetTask(0, gomock.Any()).Return(nil).MaxTimes(len(core.DutyList))

	p.AddTask(0, core.DutyList[0], queue.Task{Height: 1, Message: "a"})
	p.AddTask(0, core.DutyList[0], queue.Task{Height: 2, Message: "b"})
	p.AddTask(0, core.DutyList[0], queue.Task{Height: 3, Message: "c"})

	time.Sleep(time.Millisecond * 400)
}

func TestThreeTasksForThreeValidators(t *testing.T) {
	ctrl := gomock.NewController(t)
	q := NewMockQueue(ctrl)
	w := NewMockWork(ctrl)
	p := NewProcessor(q, w)
	p.setWorkDuration(time.Millisecond * 300)
	defer p.Close()
	defer ctrl.Finish()

	//Expectations
	w.EXPECT().Do("a").Times(1)
	w.EXPECT().Do("b").Times(1)
	w.EXPECT().Do("c").Times(1)
	q.EXPECT().GetTask(0, core.DutyList[0]).Return(nil).Times(1)
	q.EXPECT().GetTask(1, core.DutyList[0]).Return(nil).Times(1)
	q.EXPECT().GetTask(2, core.DutyList[0]).Return(nil).Times(1)

	p.AddTask(0, core.DutyList[0], queue.Task{Height: 1, Message: "a"})
	p.AddTask(1, core.DutyList[0], queue.Task{Height: 2, Message: "b"})
	p.AddTask(2, core.DutyList[0], queue.Task{Height: 3, Message: "c"})

	time.Sleep(time.Millisecond * 400)
}

func TestThreeTasksForOneValidatorAndThreeDuties(t *testing.T) {
	ctrl := gomock.NewController(t)
	q := NewMockQueue(ctrl)
	w := NewMockWork(ctrl)
	p := NewProcessor(q, w)
	p.setWorkDuration(time.Millisecond * 300)
	defer p.Close()
	defer ctrl.Finish()

	//Expectations
	w.EXPECT().Do("a").Times(1)
	w.EXPECT().Do("b").Times(1)
	w.EXPECT().Do("c").Times(1)
	q.EXPECT().GetTask(0, core.DutyList[0]).Return(nil).Times(1)
	q.EXPECT().GetTask(0, core.DutyList[1]).Return(nil).Times(1)
	q.EXPECT().GetTask(0, core.DutyList[2]).Return(nil).Times(1)

	p.AddTask(0, core.DutyList[0], queue.Task{Height: 1, Message: "a"})
	p.AddTask(0, core.DutyList[1], queue.Task{Height: 2, Message: "b"})
	p.AddTask(0, core.DutyList[2], queue.Task{Height: 3, Message: "c"})

	time.Sleep(time.Millisecond * 400)
}
