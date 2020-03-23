package sourcing

type EventSource interface {
	// Applies an event to the event source which will get recorded and applied.
	// The later means the state of the event source is updated according to the event.
	Apply(event Event)

	// Get all events that happened since the last accept, or all if accept was never called.
	Events() []Event

	Accept()
}

type eventSource struct {
	applier  EventRouter
	recorder EventRecorder
}

func newEventSource(applier EventRouter, recorder EventRecorder) *eventSource {
	return &eventSource{
		applier:  applier,
		recorder: recorder,
	}
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
