package sourcing

import ()

type EventSource interface {
	// Gets the id of the event source
	Id() EventSourceId

	// Applies an event to the event source which will get recorded and applied.
	// The later means the state of the event source is updated according to the event.
	Apply(event Event)

	// Get all events that happened since the last accept, or all if accept was never called.
	Events() []Event

	Accept()
}

type eventSource struct {
	id       EventSourceId
	applier  EventRouter
	recorder EventRecorder
}

func newEventSource(id EventSourceId, applier EventRouter, eventRecorder EventRecorder) *eventSource {
	return &eventSource{
		id:       id,
		applier:  applier,
		recorder: eventRecorder,
	}
}

func (source *eventSource) Id() EventSourceId {
	return source.id
}

func (source *eventSource) Events() []Event {
	return source.recorder.GetEvents()
}

func (source *eventSource) Apply(event Event) {
	source.applier.Route(event)
	source.recorder.Record(event)
}

func (source *eventSource) Accept() {
	source.recorder.Clear()
}
