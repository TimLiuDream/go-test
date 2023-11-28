package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Listener[T any] func(T)

type Manager[T any] interface {
	Add(n string, l Listener[T])
	Run()
}

type BaseManager[T any] struct {
	lst map[string][]Listener[T]
}

func (m *BaseManager[T]) Invoke(n string, args T) {
	for _, ls := range m.lst[n] {
		ls(args)
	}
}

func (m *BaseManager[T]) Add(n string, l Listener[T]) {
	m.lst[n] = append(m.lst[n], l)
}

type Command struct {
	Kind string
	Args []string
}

type CommandEventManager struct {
	BaseManager[*Command]
}

func (m *CommandEventManager) Run() {
	var (
		inp  string
		args Command
	)

	fmt.Scanln(&inp)

	cmd := strings.Split(inp, ":")

	if l := len(cmd); l == 0 {
		m.Invoke("no-command", nil)
	} else if l > 1 {
		args.Args = strings.Split(cmd[1], " ")
	}

	args.Kind = cmd[0]

	m.Invoke("any-command", &args)
	m.Invoke(args.Kind, &args)
}

func NewCommandEventManager() Manager[*Command] {
	return &CommandEventManager{
		BaseManager: BaseManager[*Command]{lst: make(map[string][]Listener[*Command])},
	}
}

func main() {
	cem := NewCommandEventManager()

	cem.Add("no-command", func(_ *Command) {
		fmt.Println("no command was recieved")
	})

	cem.Add("any-command", func(c *Command) {
		fmt.Printf("the %s command was executed", c.Kind)
	})

	cem.Add("sum", func(c *Command) {
		a, _ := strconv.Atoi(c.Args[0])
		b, _ := strconv.Atoi(c.Args[1])
		fmt.Printf("the sum result is: %d", a+b)
	})

	cem.Run()
}
