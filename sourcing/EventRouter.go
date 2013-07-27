package sourcing

type EventRouter interface {
	Route(e Event)
}
