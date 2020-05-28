package main

import (
	"fmt"
	"log"
)

func main() {
	//m := newMessage("Welcome to Builder Gophers")
	//g := newGreeter(m)
	//e := newEvent(g)
	msg := "Welcome to Builder Gophers"
	e, err := initializeEvent(msg)
	if err != nil {
		log.Fatalf("FAILURE: %v", err)
	}
	e.Start()
}

func newMessage(msg string) message {
	return message{value: msg}
}

type message struct {
	value string
}

func (m message) String() string {
	return fmt.Sprintf("[OFFICIAL MESSAGE]: %s", m.value)
}

func newGreeter(m message) greeter {
	return greeter{message: m}
}

type greeter struct {
	message message
}

func (g greeter) greet() string {
	return g.message.String()
}

func newEvent(g greeter) (event, error) {
	//return event{greeter: g}
	//return event{greeter: g}, errors.New("failed event")
	return event{greeter: g}, nil
}

type event struct {
	greeter greeter
}

func (e event) Start() {
	fmt.Println(e.greeter.greet())
}
