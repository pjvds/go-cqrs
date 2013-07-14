package eventing

import (
	"reflect"
)

type EventNamer interface {
	GetEventName(e Event) EventName
	GetEventNameFromType(eventType reflect.Type) EventName
}

func NewTypeEventNamer() *TypeEventNamer {
	return &TypeEventNamer{}
}

type TypeEventNamer struct {
}

func (namer *TypeEventNamer) GetEventName(e Event) EventName {
	t := reflect.TypeOf(e)
	return namer.GetEventNameFromType(t)
}

func (namer *TypeEventNamer) GetEventNameFromType(t reflect.Type) EventName {
	return EventName(t.PkgPath() + "/" + t.Name())
}
