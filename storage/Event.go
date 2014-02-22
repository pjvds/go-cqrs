package storage

import (
	"fmt"
	"github.com/dominikmayer/go-cqrs/sourcing"
	"time"
)

// Holds the meta information for an event.
type Event struct {
	EventId   EventId        `json:"eventId"`   // The id of the event itself
	Name      EventName      `json:"name"`      // The event name, this value is also used for type identification (maps name to Go type)
	Sequence  EventSequence  `json:"sequence"`  // The sequence of the event which starts at zero.
	Timestamp time.Time      `json:"timestamp"` // The point in time when this event happened.
	Data      sourcing.Event `json:"payload"`   // The data of the event.
}

func NewEvent(eventId EventId, name EventName, sequence EventSequence, timestamp time.Time, data sourcing.Event) *Event {
	return &Event{
		EventId:   eventId,
		Name:      name,
		Sequence:  sequence,
		Timestamp: timestamp,
		Data:      data,
	}
}

// Returns a string representation of the EventEnvelope in the "{sequence}/{eventname}" format.
func (e *Event) String() string {
	return fmt.Sprintf("%v/%v", e.Sequence, e.Name)
}
