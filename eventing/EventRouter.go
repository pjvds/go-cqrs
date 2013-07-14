package eventing

type EventRouter interface {
	RegisterHandler(eventName EventName, handler EventHandler)
	Route(e EventEnvelope)
}

func NewEventRouter() EventRouter {
	return &eventRouter{
		handlers: make(map[EventName]EventHandler),
	}
}

type eventRouter struct {
	handlers map[EventName]EventHandler
}

func (e *eventRouter) RegisterHandler(eventName EventName, handler EventHandler) {
	e.handlers[eventName] = handler
}

func (router *eventRouter) Route(e EventEnvelope) {
	if handler := router.handlers[e.Name]; handler != nil {
		handler(e.Payload)
	}
}
