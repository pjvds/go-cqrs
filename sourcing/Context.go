package sourcing

import ()

func newDefaultContext() *Context {
	return &Context{}
}

// The sourcing context.
type Context struct {
}

func (ctx *Context) create(id EventSourceId, source interface{}) EventSource {
	recorder := NewEventRecorder()
	router := NewReflectBasedRouter(source)

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
