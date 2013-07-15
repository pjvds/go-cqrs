package sourcing

import (
	"github.com/nu7hatch/gouuid"
)

type EventSourceId uuid.UUID

func NewEventSourceId() EventSourceId {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return EventSourceId(*guid)
}
