package sourcing

type Version int64

type EventSource interface {
	GetVersion() Version

	// Applies an event to the event source which will get recorded and applied.
	// The later means the state of the event source is updated according to the event.
	Apply(event Event)

	// Get all events that happened since the last accept, or all if accept was never called.
	Events() []Event

	Accept(version Version)
}

type eventSource struct {
	version  Version
	applier  EventRouter
	recorder EventRecorder
}

func newEventSource(version Version, applier EventRouter, recorder EventRecorder) *eventSource {
	return &eventSource{
		version:  version,
		applier:  applier,
		recorder: recorder,
	}
}

func (source *eventSource) GetVersion() Version {
	return source.version
}

func (source *eventSource) Events() []Event {
	return source.recorder.GetEvents()
}

func (source *eventSource) Apply(event Event) {
	source.applier.Route(event)
	source.recorder.Record(event)
}

func (source *eventSource) Accept(version Version) {
	source.recorder.Clear()
	source.version = version
}
