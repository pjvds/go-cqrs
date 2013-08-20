package sourcing

import ()

func newDefaultContext() *Context {
	return &Context{
		newRecorder: func(id EventSourceId, source interface{}) EventRecorder {
			return NewEventRecorder()
		},
		newRouterForSource: func(id EventSourceId, source interface{}) EventRouter {
			return NewReflectBasedRouter(source)
		},
	}
}

// The sourcing context.
type Context struct {
	newRecorder        func(id EventSourceId, source interface{}) EventRecorder
	newRouterForSource func(id EventSourceId, source interface{}) EventRouter
}

func (ctx *Context) create(id EventSourceId, source interface{}) EventSource {
	recorder := ctx.newRecorder(id, source)
	router := ctx.newRouterForSource(id, source)

	return newEventSource(id, router, recorder)
}

// Creates a new EventSource object that can be used to source events.
func (ctx *Context) CreateNew(source interface{}) EventSource {
	return ctx.create(NewEventSourceId(), source)
}

// Creates an existing EventSource object based on the state from the history
// are replays history so the specified source can update it's state.
func (ctx *Context) CreateFromHistory(source interface{}, id EventSourceId, history []Event) EventSource {
	eventSource := ctx.create(id, source)

	for _, event := range history {
		eventSource.Apply(event)
	}

	return eventSource
}
