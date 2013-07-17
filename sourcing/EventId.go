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
