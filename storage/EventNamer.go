package storage

import (
	"github.com/pjvds/go-cqrs/sourcing"
	"reflect"
)

type EventNamer interface {
	GetEventName(e sourcing.Event) EventName
	GetEventNameFromType(eventType reflect.Type) EventName
}

func NewTypeEventNamer() *TypeEventNamer {
	return &TypeEventNamer{}
}

type TypeEventNamer struct {
}

func (namer *TypeEventNamer) GetEventName(e sourcing.Event) EventName {
	t := reflect.TypeOf(e)
	return namer.GetEventNameFromType(t)
}

func (namer *TypeEventNamer) GetEventNameFromType(t reflect.Type) EventName {
	return EventName(t.PkgPath() + "/" + t.Name())
}
