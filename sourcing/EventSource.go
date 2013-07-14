package sourcing

type EventSourceState interface {
	Events() []EventEnvelope
}

type EventSource interface {
	Apply(event Event)
}

type eventSource struct {
	eventNamer EventNamer
	events     []EventEnvelope
	router     EventRouter
}

func newEventSource(eventNamer EventNamer, router EventRouter) *eventSource {
	return &eventSource{
		eventNamer: eventNamer,
		router:     router,
	}
}

func (source *eventSource) Events() []EventEnvelope {
	return source.events
}

func (source *eventSource) Apply(event Event) {
	envelope := EventEnvelope{
		Name:    source.eventNamer.GetEventName(event),
		Payload: event,
	}
	source.router.Route(envelope)
	source.events = append(source.events, envelope)
}
