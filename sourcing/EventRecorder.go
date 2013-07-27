package sourcing

type EventRecorder struct {
	events []*Event
}

func NewEventRecorder() *EventRecorder {
	return &EventRecorder{
		events: make([]*Event, 0),
	}
}

func (r *EventRecorder) Record(e *Event) {
	r.events = append(r.events, e)
}
