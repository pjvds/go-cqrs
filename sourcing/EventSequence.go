package sourcing

type EventSequence int64

func NewEventSequence(value int64) EventSequence {
	return EventSequence(value)
}
