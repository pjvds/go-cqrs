package sourcing

import ()

func newDefaultContext() *Context {
	return &Context{
		sources: make(map[interface{}]*eventSource, 5),
	}
}

// The sourcing context. Objects can be attached into this context so they
// monitored for change. Do not forget to detach them when they are not needed
// anymore.
type Context struct {
	sources map[interface{}]*eventSource
}

func (ctx *Context) attach(id EventSourceId, source interface{}) EventSource {
	router := NewReflectBasedRouter(source)

	eventSource := newEventSource(id, router)

	ctx.sources[source] = eventSource
	return eventSource
}

// Creates a new EventSource object that can be used to source events.
func (ctx *Context) AttachNew(source interface{}) EventSource {
	return ctx.attach(NewEventSourceId(), source)
}

// Creates an existing EventSource object based on the state from the history
// are replays history so the specified source can update it's state.
func (ctx *Context) AttachFromHistory(source interface{}, id EventSourceId, history []Event) EventSource {
	eventSource := ctx.attach(id, source)

	for _, event := range history {
		eventSource.Apply(event)
	}

	return eventSource
}

func (ctx *Context) GetState(source interface{}) EventSource {
	state, ok := ctx.sources[source]
	if ok {
		return state
	} else {
		return nil
	}
}

// Removes the source from the context. It releases all references to the source
// and the related EventSource.
// A source should not generate any events after it is detached.
func (ctx *Context) Detach(source interface{}) {
	delete(ctx.sources, source)
}
