package storage

import (
	"github.com/pjvds/go-cqrs/sourcing"
	"reflect"
)

// See: http://play.golang.org/p/7F__is92pX

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
	Log.Notice("Getting name from: %T", e)

	t := reflect.TypeOf(e)
	return namer.GetEventNameFromType(t)
}

func (namer *TypeEventNamer) GetEventNameFromType(t reflect.Type) EventName {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	Log.Notice("Getting name from type: %v", t.String())
	return EventName(t.PkgPath() + "/" + t.Name())
}
