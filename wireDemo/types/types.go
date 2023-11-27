package types

import (
	"errors"
	"fmt"
	"time"
)

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter struct {
	Message
	Grumpy bool
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return "go away!"
	}
	return g.Message
}

type Event struct {
	Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewEventNumber() int {
	return 1
}
