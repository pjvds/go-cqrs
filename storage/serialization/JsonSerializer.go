package serialization

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pjvds/go-cqrs/storage"
	"reflect"
	"time"
)

type JsonDocument map[string]interface{}

// Holds the meta information for an event with the data
// object as a raw message so that it can be deserialized in two stages.
type jsonEvent struct {
	EventId   storage.EventId       `json:"eventId"`   // The id of the event itself
	Name      storage.EventName     `json:"name"`      // The event name, this value is also used for type identification (maps name to Go type)
	Sequence  storage.EventSequence `json:"sequence"`  // The sequence of the event which starts at zero.
	Timestamp time.Time             `json:"timestamp"` // The point in time when this event happened.
	Data      json.RawMessage       `json:"payload"`   // The data of the event.
}

type JsonSerializer struct {
	types *EventTypeRegister
}

func NewJsonSerializer(types *EventTypeRegister) *JsonSerializer {
	return &JsonSerializer{
		types: types,
	}
}

func (s *JsonSerializer) Serialize(e *storage.Event) ([]byte, error) {
	data, err := json.Marshal(e)

	return data, err
}

func (s *JsonSerializer) Deserialize(name storage.EventName, data []byte) (*storage.Event, error) {
	raw := new(jsonEvent)

	if err := json.Unmarshal(data, raw); err != nil {
		return nil, err
	}

	e := new(storage.Event)
	e.EventId = raw.EventId
	e.Name = raw.Name
	e.Sequence = raw.Sequence
	e.Timestamp = raw.Timestamp

	eventType, ok := s.types.Get(name)
	if !ok {
		return e, errors.New(fmt.Sprintf("No known type for %v, register it first", name.String()))
	}
	eventValue := reflect.New(eventType)
	event := eventValue.Interface()
	if err := json.Unmarshal(raw.Data, event); err != nil {
		return nil, err
	}

	e.Data = event

	return e, nil
}
