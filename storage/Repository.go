package storage

import (
	"encoding/json"
	"github.com/pjvds/go-cqrs/sourcing"
	"time"
)

type RepositoryBackend interface {
	WriteStream(change *EventStreamChange) error
}

type Repository struct {
	namer   EventNamer
	backend RepositoryBackend
}

func NewRepository(backend RepositoryBackend) *Repository {
	return &Repository{
		namer:   NewTypeEventNamer(),
		backend: backend,
	}
}

func (r *Repository) Add(source sourcing.EventSource) error {
	change, err := r.getStreamChangeFromSource(source)
	if err != nil {
		return err
	}

	return r.backend.WriteStream(change)
}

func (r *Repository) getStreamChangeFromSource(source sourcing.EventSource) (*EventStreamChange, error) {
	events := source.Events()
	eventCount := int64(len(events))

	fromSequence := int64(0)
	streamEvents := make([]*Event, eventCount)

	for i, e := range events {
		data, err := json.Marshal(e)
		if err != nil {
			return nil, err
		}

		name := r.namer.GetEventName(e)
		sequence := NewEventSequence(fromSequence + int64(i))
		streamEvents[i] = NewEvent(NewEventId(), name, sequence, time.Now(), data)
	}

	return &EventStreamChange{
		StreamId: EventStreamId(source.Id()),
		From:     fromSequence,
		To:       fromSequence + eventCount,
		Events:   streamEvents,
	}, nil
}
