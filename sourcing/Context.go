package sourcing

import ()

func newDefaultContext() *Context {
	return &Context{
		sources: make(map[EventSource]*SourceState, 5),
		namer:   NewTypeEventNamer(),
	}
}

type Context struct {
	sources map[EventSource]*SourceState
	namer   EventNamer
}

type SourceState struct {
	Events  []EventEnvelope
	applier EventHandler
	router  EventRouter
}

func (ctx *Context) AttachNew(source EventSource) {
	Log.Debug("Attaching new %T", source)

	router := NewReflectBasedRouter(ctx.namer, source)
	state := &SourceState{
		Events: make([]EventEnvelope, 0),
		router: router,
	}
	state.applier = func(e Event) {
		name := ctx.namer.GetEventName(e)
		event := e

		envelop := EventEnvelope{
			Name:    name,
			Payload: event,
		}

		Log.Debug("Applying %v to %T\n\tstate: %+v", envelop.Name, source, event)
		router.Route(envelop)
		state.Events = append(state.Events, envelop)
	}
	source.SetEventApplier(state.applier)
	ctx.sources[source] = state
}

func (ctx *Context) AttachWithHistory(source EventSource, history []EventEnvelope) {
	Log.Debug("Attaching new %T", source)

	router := NewReflectBasedRouter(ctx.namer, source)
	state := &SourceState{
		Events: make([]EventEnvelope, 0),
		router: router,
	}
	state.applier = func(e Event) {
		name := ctx.namer.GetEventName(e)
		event := e

		envelop := EventEnvelope{
			Name:    name,
			Payload: event,
		}

		Log.Debug("Applying %v to %T\n\tstate: %+v", envelop.Name, source, event)
		router.Route(envelop)
		state.Events = append(state.Events, envelop)
	}
	source.SetEventApplier(state.applier)
	ctx.sources[source] = state

	for _, e := range history {
		router.Route(e)
	}
}

func (ctx *Context) GetState(source EventSource) *SourceState {
	state, ok := ctx.sources[source]
	if ok {
		return state
	} else {
		return nil
	}
}

func (ctx *Context) Detach(source EventSource) {
	delete(ctx.sources, source)
}
