package storage

import (
	"github.com/pjvds/gouuid"
)

type EventStreamId uuid.UUID

func NewEventStreamId() EventStreamId {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return EventStreamId(*guid)
}

func (id EventStreamId) String() string {
	guid := uuid.UUID(id)
	return guid.String()
}

func (id EventStreamId) MarshalJSON() ([]byte, error) {
	value := uuid.UUID(id)
	return value.MarshalJSON()
}

func (id *EventStreamId) UnmarshalJSON(b []byte) error {
	value := uuid.UUID(*id)
	err := value.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	eventId := EventStreamId(value)
	id = &eventId
	return nil
}
