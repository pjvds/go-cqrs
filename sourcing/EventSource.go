package sourcing

import ()

type EventSource interface {
	Id() EventSourceId
	Apply(event Event)
	Events() []Event
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
