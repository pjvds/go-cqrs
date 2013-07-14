package eventing

// Represents the unique identifiable name of an event type within a context.
type EventName string

func (name EventName) String() string {
	return string(name)
}
