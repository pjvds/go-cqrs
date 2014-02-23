package storage

type EventStreamChange struct {
	StreamId EventStreamId
	From     int64
	To       int64
	Events   []*Event
}

type EventStreamChangePersist struct {
	StreamId string
	From     int64
	To       int64
	Events   []*Event
}

func (e *EventStreamChange) GetPersistableObject() *EventStreamChangePersist {
	return &EventStreamChangePersist{
		StreamId:  e.StreamId.String(),
		From: e.From,
		To: e.To,
		Events: e.Events,
	}
}
