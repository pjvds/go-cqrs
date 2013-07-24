package sourcing

import (
	"github.com/nu7hatch/gouuid"
)

type EventId uuid.UUID

func NewEventId() EventId {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return EventId(*guid)
}

func (id EventId) String() string {
	guid := uuid.UUID(id)
	return guid.String()
}

func (id EventId) MarshalJSON() ([]byte, error) {
	value := uuid.UUID(id)
	return value.MarshalJSON()
}

func (id *EventId) UnmarshalJSON(b []byte) error {
	value := uuid.UUID(*id)
	err := value.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	eventId := EventId(value)
	id = &eventId
	return nil
}
