package storage

type EventStreamChange struct {
	StreamId EventStreamId
	From     int64
	To       int64
	Events   []*Event
}
