package sourcing

import ()

func newDefaultContext() *Context {
	return &Context{
		sources: make(map[interface{}]*eventSource, 5),
		namer:   NewTypeEventNamer(),
	}
}

type Context struct {
	sources map[interface{}]*eventSource
	namer   EventNamer
}

type SourceState struct {
	Events  []EventEnvelope
	applier EventHandler
	router  EventRouter
}

func (ctx *Context) attach(id EventSourceId, source interface{}) EventSource {
	namer := ctx.namer
	router := NewReflectBasedRouter(ctx.namer, source)

	eventSource := newEventSource(id, namer, router)

	ctx.sources[source] = eventSource
	return eventSource
}

func (ctx *Context) AttachNew(source interface{}) EventSource {
	return ctx.attach(NewEventSourceId(), source)
}

func (ctx *Context) AttachFromHistory(source interface{}, history []EventEnvelope) EventSource {
	id := history[0].EventSourceId
	eventSource := ctx.attach(id, source)

	for _, event := range history {
		eventSource.Apply(event.Payload)
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

func (ctx *Context) Detach(source interface{}) {
	delete(ctx.sources, source)
}
