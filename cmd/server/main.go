package main

import (
	"ssv/internal"
	"ssv/internal/processor"
	"ssv/internal/queue"
)

type Worker struct {
}

func (w Worker) Do(message string) {
	println(message)
}

// This example demonstrates a trivial echo server.
func main() {
	w := Worker{}

	queue := queue.NewQueue()
	p := processor.NewProcessor(queue, w)
	app := internal.NewApp(p)
	app.Run()
}
