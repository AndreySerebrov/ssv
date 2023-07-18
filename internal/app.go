package internal

import (
	"io"
	"net/http"
	"ssv/internal/queue"

	"golang.org/x/net/websocket"
)

type Duty struct {
	Validator int    `json:"validator"`
	Duty      string `json:"duty"`
	Height    int    `json:"height"`
}

type Error struct {
	Cause string `json:"cause"`
}

type Processor interface {
	AddTask(validator int, dutyType string, task queue.Task)
}

type App struct {
	processor Processor
}

func NewApp(p Processor) *App {
	return &App{
		processor: p,
	}
}

// Echo the data received on the WebSocket.
func (a App) Server(ws *websocket.Conn) {
	duty := Duty{}
	err := websocket.JSON.Receive(ws, &duty)
	if err != nil {
		panic(err)
	}
	task := queue.Task{
		Height: duty.Height,
	}
	a.processor.AddTask(duty.Validator, duty.Duty, task)
}

func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func (a *App) Run() {
	http.Handle("/", websocket.Handler(a.Server))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	// http.Handle("/echo", websocket.Handler(EchoServer))
	// err := http.ListenAndServe(":5000", nil)
	// if err != nil {
	// 	panic("ListenAndServe: " + err.Error())
	// }
}
