package storage

import (
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
	Log.Debug("Adding changes for %v", source.Id())

	change, err := r.getStreamChangeFromSource(source)
	if err != nil {
		return err
	}

	Log.Debug("Saving change with backend")
	if err := r.backend.WriteStream(change); err != nil {
		Log.Error("Error from backend: %v", err)
		return err
	}
	Log.Debug("Succesfully wrote change to repository backend.")

	Log.Debug("Starting dispatching change.")
	r.dispatcher.Dispatch(change)

	Log.Debug("Change added succesfully")
	return nil
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
		Log.Notice("%T", e)

		name := r.namer.GetEventName(e)
		sequence := NewEventSequence(fromSequence + int64(i))

		id := NewEventId()
		timestamp := time.Now()
		streamEvents[i] = NewEvent(id, name, sequence, timestamp, e)
	}

	return &EventStreamChange{
		StreamId: EventStreamId(sourceId),
		From:     fromSequence,
		To:       fromSequence + eventCount,
		Events:   streamEvents,
	}, nil
}
