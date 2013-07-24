package sourcing

type EventSequence int64

func NewEventSequence(value int64) EventSequence {
	return EventSequence(value)
}

func (sequence EventSequence) Next() EventSequence {
	value := int64(sequence)
	return NewEventSequence(value + 1)
}
