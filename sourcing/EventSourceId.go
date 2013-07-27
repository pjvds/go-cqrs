package sourcing

import (
	"github.com/pjvds/gouuid"
)

type EventSourceId uuid.UUID

func NewEventSourceId() EventSourceId {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return EventSourceId(*guid)
}

func ParseEventSourceId(value string) (id EventSourceId, err error) {
	guid := new(uuid.UUID)
	if guid, err = uuid.ParseHex(value); err == nil {
		id = EventSourceId(*guid)
	}

	return
}

func (id EventSourceId) String() string {
	guid := uuid.UUID(id)
	return guid.String()
}

func (id EventSourceId) MarshalJSON() ([]byte, error) {
	value := uuid.UUID(id)
	return value.MarshalJSON()
}

func (id *EventSourceId) UnmarshalJSON(b []byte) error {
	value := uuid.UUID(*id)
	err := value.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	*id = EventSourceId(value)
	return nil
}
