package sourcing

import (
	"reflect"
)

type EventTypeRegister struct {
	types map[EventName]reflect.Type
}

func NewEventTypeRegister() *EventTypeRegister {
	return &EventTypeRegister{
		types: make(map[EventName]reflect.Type, 0),
	}
}

func (register *EventTypeRegister) Register(n EventName, t reflect.Type) {
	register.types[n] = t
}

func (register *EventTypeRegister) Get(n EventName) (reflect.Type, bool) {
	t, ok := register.types[n]
	return t, ok
}
