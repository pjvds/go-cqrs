package sourcing

import (
	"time"
)

type EventSource interface {
	Id() EventSourceId
	Apply(event Event)
	Events() []EventEnvelope
}

type eventSource struct {
	id         EventSourceId
	eventNamer EventNamer
	events     []EventEnvelope
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

func (source *eventSource) Events() []EventEnvelope {
	return source.events
}

func (source *eventSource) Apply(event Event) {
	envelope := EventEnvelope{
		EventSourceId: source.Id(),
		EventId:       NewEventId(),
		Name:          source.eventNamer.GetEventName(event),
		Timestamp:     time.Now(),
		Payload:       event,
	}
	source.router.Route(envelope)
	source.events = append(source.events, envelope)
}
