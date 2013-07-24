package sourcing

import (
	"fmt"
	"time"
)

// Holds the meta information for an event.
type EventEnvelope struct {
	EventSourceId EventSourceId `json:"eventSourceId"` // The id of the source that owns this event
	EventId       EventId       `json:"eventId"`       // The id of the event itself
	Name          EventName     `json:"name"`          // The event name, this value is also used for type identification (maps name to Go type)
	Sequence      EventSequence `json:"sequence"`      // The sequence of the event which starts at zero.
	Timestamp     time.Time     `json:"timestamp"`     // The point in time when this event happened.
	Payload       Event         `json:"payload"`       // The data of the event.
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

// Returns a string representation of the EventEnvelope in the "{sequence}/{eventname}" format.
func (e *EventEnvelope) String() string {
	return fmt.Sprintf("%v/%v", e.Sequence, e.Name)
}

func PackEvents(eventSourceId EventSourceId, events []Event) []*EventEnvelope {
	envelopes := make([]*EventEnvelope, len(events))

	for index, event := range events {
		envelopes[index] = NewEventEnvelope(eventSourceId, NewEventId(),
			defaultContext.namer.GetEventName(event), NewEventSequence(int64(index)),
			time.Now(), event)
	}

	return envelopes
}
