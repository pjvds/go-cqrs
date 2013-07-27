package storage

import (
	"reflect"
)

// A register that holds the mapping between an event name and it's static type.
// All event type should be registered at bootstrap time so that an event store,
// bus or other services can deserialize messages to concrete types.
type EventTypeRegister struct {
	types map[EventName]reflect.Type
}

// Creates a new register with an empty map.
func NewEventTypeRegister() *EventTypeRegister {
	return &EventTypeRegister{
		types: make(map[EventName]reflect.Type, 0),
	}
}

// Registers an event type. An existing entry with the same name is overwritten
// if it exists.
func (register *EventTypeRegister) Register(n EventName, t reflect.Type) {
	register.types[n] = t
}

// Get the static type from an event name. It results `true` for `ok` if
// the type was found; otherwise, `false`.
func (register *EventTypeRegister) Get(n EventName) (reflect.Type, bool) {
	t, ok := register.types[n]
	return t, ok
}
