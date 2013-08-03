package memory

import (
	. "github.com/pjvds/go-cqrs/storage"
)

type MemoryBackend struct {
	changes map[EventStreamId][]*EventStreamChange
}

func NewMemoryBackend() *MemoryBackend {
	return &MemoryBackend{
		changes: make(map[EventStreamId][]*EventStreamChange, 0),
	}
}

func (m *MemoryBackend) WriteStream(change *EventStreamChange) error {
	changes, ok := m.changes[change.StreamId]

	if !ok {
		changes = make([]*EventStreamChange, 0, 1)
	}

	m.changes[change.StreamId] = append(changes, change)
	return nil
}

func (m *MemoryBackend) ReadStream(streamId EventStreamId) ([]*Event, error) {
	changes, ok := m.changes[streamId]

	if !ok {
		return nil, nil
	}

	events := make([]*Event, 0, len(changes)*2)

	for _, change := range changes {
		for _, event := range change.Events {
			events = append(events, event)
		}
	}

	return events, nil
}
