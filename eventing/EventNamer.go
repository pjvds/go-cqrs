package eventing

import (
	"reflect"
)

type EventNamer interface {
	GetEventName(e Event) EventName
	GetEventNameFromType(eventType reflect.Type) EventName
}

func NewEventNamer() *eventNamer {
	return &eventNamer{}
}

type eventNamer struct {
}

func (namer *eventNamer) GetEventName(e Event) EventName {
	t := reflect.TypeOf(e)
	return namer.GetEventNameFromType(t)
}

func (namer *eventNamer) GetEventNameFromType(t reflect.Type) EventName {
	return EventName(t.String())
}
