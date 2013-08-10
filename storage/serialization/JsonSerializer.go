package serialization

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pjvds/go-cqrs/storage"
	"reflect"
	"time"
)

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
	raw := new(struct {
		EventId   storage.EventId       `json:"eventId"`
		Name      storage.EventName     `json:"name"`
		Sequence  storage.EventSequence `json:"sequence"`
		Timestamp time.Time             `json:"timestamp"`
		Data      json.RawMessage       `json:"payload"`
	})

	if err := json.Unmarshal(data, raw); err != nil {
		Log.Error("Error while unmarhalling data: %v\n\tdata: \"%v\"", err.Error(), string(data))
		return nil, fmt.Errorf("Unable to unmarshall data: %v", err)
	}

	e := new(storage.Event)
	e.EventId = raw.EventId
	e.Name = raw.Name
	e.Sequence = raw.Sequence
	e.Timestamp = raw.Timestamp

	eventType, ok := s.types.Get(e.Name)
	if !ok {
		return e, errors.New(fmt.Sprintf("No known type for '%v', register it first", e.Name))
	}
	eventValue := reflect.New(eventType)
	event := eventValue.Interface()
	if err := json.Unmarshal(raw.Data, event); err != nil {
		return nil, err
	}

	e.Data = event

	return e, nil
}
