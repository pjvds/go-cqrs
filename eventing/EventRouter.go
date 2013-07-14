package eventing

type EventRouter interface {
	Route(e EventEnvelope)
}
