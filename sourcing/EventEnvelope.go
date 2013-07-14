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

func PackEvents(events []Event) []EventEnvelope {
	envelopes := make([]EventEnvelope, len(events))

	for index, event := range events {
		envelopes[index] = EventEnvelope{
			Name:    defaultContext.namer.GetEventName(event),
			Payload: event,
		}
	}

	return envelopes
}
