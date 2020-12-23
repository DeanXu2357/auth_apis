package event_listener

import (
	"log"
)

type Dispatcher struct {
	listeners map[string][]Listener
	tasks chan Event
}

type Listener interface {
	Handle(Event) error
}

type Event interface {
	GetName() string
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{tasks: make(chan Event), listeners: make(map[string][]Listener, 0)}
}

func (d *Dispatcher) Fake() {
	d.listeners = map[string][]Listener{}
}

func (d *Dispatcher) Dispatch(e Event) {
	d.tasks <- e
}

func (d *Dispatcher) Consume() {
	go func() {
		for t := range d.tasks {
			eventName := t.GetName()
			execute(d.listeners[eventName], t)
		}
	}()
}

func execute(listeners []Listener, e Event) {
	for _, listener := range listeners {
		if err := listener.Handle(e); err != nil {
			log.Println(err)
		}
	}
}

func (d *Dispatcher) AttachListener(eventName string, listener Listener) {
	d.listeners[eventName] = append(d.listeners[eventName], listener)
}

func (d *Dispatcher) Close() {
	close(d.tasks)
}
