package main

import (
	"fmt"
)

type Message string

func NewMessage() Message {
	return Message("Hello, world")
}

type Greeter struct {
	Message Message
}

func NewGreater(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// with wire, we use dependency injection like this
func main() {
	e := InitializeEvent()
	e.Start()
}

// with wire, we use dependency injection like this
// func main() {
// 	//这里的初始化步骤太多了
// 	message := NewMessage()
// 	greater := NewGreater(message)
// 	event := NewEvent(greater)

// 	event.Start()
// }
