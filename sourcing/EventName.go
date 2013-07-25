package sourcing

import ()

// Represents the unique identifiable name of an event type within a context.
type EventName string

func NewEventName(value string) *EventName {
	name := EventName(value)
	return &name
}

func (name EventName) String() string {
	return string(name)
}
