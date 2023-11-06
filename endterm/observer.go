package main

import "fmt"

type Event struct {
	Desciption string
}

type Observer interface {
	OnNotify(Event)
}

type EventNotifier struct {
	observers []Observer
}

func (e *EventNotifier) Notify(event Event) {
	for _, observer := range e.observers {
		observer.OnNotify(event)
	}
}

func (e *EventNotifier) RegisterObserver(o Observer) {
	e.observers = append(e.observers, o)
}

func (c *Character) OnNotify(event Event) {
	fmt.Printf("%s", event.Desciption)
}
