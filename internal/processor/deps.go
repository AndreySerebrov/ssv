package processor

import "ssv/internal/queue"

//go:generate mockgen -source=deps.go -destination=./mocks.go -package=processor

type Queue interface {
	AddTask(validator int, taskType string, task queue.Task)
	GetTask(validator int, taskType string) *queue.Task
}

type Work interface {
	Do(message string)
}
