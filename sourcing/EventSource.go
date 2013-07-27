package sourcing

import ()

type EventSource interface {
	Id() EventSourceId
	Apply(event Event)
	Events() []Event
}

type eventSource struct {
	id         EventSourceId
	eventNamer EventNamer
	events     []Event
	router     EventRouter
}

func newEventSource(id EventSourceId, eventNamer EventNamer, router EventRouter) *eventSource {
	return &eventSource{
		id:         id,
		eventNamer: eventNamer,
		router:     router,
	}
}

func (source *eventSource) Id() EventSourceId {
	return source.id
}

func (source *eventSource) Events() []Event {
	return source.events
}

func (source *eventSource) Apply(event Event) {
	source.router.Route(event)
	source.events = append(source.events, event)
}
