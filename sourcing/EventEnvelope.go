package sourcing

import ()

// Holds the meta information for an event.
type EventEnvelope struct {
	Name    EventName
	Payload Event
}

func (e *EventEnvelope) String() string {
	return e.Name.String()
}
