package sourcing

type EventRecorder interface {
	// Records the given event.
	Record(e Event)

	// Gets the recorded events, or an empty slice if none.
	GetEvents() []Event

	// Clears the recorded events.
	Clear()
}

type eventRecorder struct {
	events []Event
}

func NewEventRecorder() EventRecorder {
	return &eventRecorder{
		events: make([]Event, 0),
	}
}

func (r *eventRecorder) Record(e Event) {
	r.events = append(r.events, e)
}

func (r *eventRecorder) GetEvents() []Event {
	return r.events
}

func (r *eventRecorder) Clear() {
	r.events = make([]Event, 0)
}
