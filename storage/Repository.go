package storage

import (
	"encoding/json"
	"github.com/pjvds/go-cqrs/sourcing"
	"time"
)

type EventDispatcher interface {
	Dispatch(change *EventStreamChange)
}

type NullEventDispatcher struct{}

func (dispatcher *NullEventDispatcher) Dispatch(change *EventStreamChange) {
}
func NewNullEventDispatcher() *NullEventDispatcher {
	return &NullEventDispatcher{}
}

type RepositoryBackend interface {
	WriteStream(change *EventStreamChange) error
	ReadStream(streamId EventStreamId) ([]*Event, error)
}

type Repository struct {
	namer      EventNamer
	backend    RepositoryBackend
	dispatcher EventDispatcher
}

func NewRepository(backend RepositoryBackend, dispatcher EventDispatcher) *Repository {
	return &Repository{
		namer:      NewTypeEventNamer(),
		backend:    backend,
		dispatcher: dispatcher,
	}
}

func (r *Repository) Add(source sourcing.EventSource) error {
	change, err := r.getStreamChangeFromSource(source)
	if err != nil {
		return err
	}

	if err = r.backend.WriteStream(change); err != nil {
		r.dispatcher.Dispatch(change)
	}

	return err
}

func (r *Repository) Get(sourceId sourcing.EventSourceId, source sourcing.EventSource) error {
	events, err := r.backend.ReadStream(EventStreamId(sourceId))
	if err != nil {
		return err
	}

	for _, e := range events {
		source.Apply(e)
	}

	return nil
}

func (r *Repository) getStreamChangeFromSource(source sourcing.EventSource) (*EventStreamChange, error) {
	sourceId := source.Id()
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
		StreamId: EventStreamId(sourceId),
		From:     fromSequence,
		To:       fromSequence + eventCount,
		Events:   streamEvents,
	}, nil
}
