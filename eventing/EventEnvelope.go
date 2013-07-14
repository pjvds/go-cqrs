package eventing

// Holds the meta information for an event.
type EventEnvelope struct {
	Name    EventName
	Payload Event
}
