package sourcing

import (
	"time"
)

// Holds the meta information for an event.
type EventEnvelope struct {
	EventSourceId EventSourceId
	EventId       EventId
	Name          EventName
	Sequence      EventSequence
	Timestamp     time.Time
	Payload       Event
}

func NewEventEnvelope(eventSourceId EventSourceId, eventId EventId, name EventName, sequence EventSequence, timestamp time.Time, payload Event) *EventEnvelope {
	return &EventEnvelope{
		EventSourceId: eventSourceId,
		EventId:       eventId,
		Name:          name,
		Sequence:      sequence,
		Timestamp:     timestamp,
		Payload:       payload,
	}
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
