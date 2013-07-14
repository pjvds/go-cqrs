package sourcing

type EventRouter interface {
	Route(e EventEnvelope)
}
