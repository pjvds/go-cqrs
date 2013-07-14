package eventing

import ()

var (
	DefaultContext = newDefaultContext()
)

func newDefaultContext() *Context {
	return &Context{
		sources: make(map[EventSource]*SourceState, 5),
		namer:   NewEventNamer(),
	}
}

type Context struct {
	sources map[EventSource]*SourceState
	namer   EventNamer
}

type SourceState struct {
	Events  []*EventEnvelope
	applier EventHandler
}

func (ctx *Context) Attach(source EventSource) {
	Log.Debug("Attaching %T", source)

	state := &SourceState{
		Events: make([]*EventEnvelope, 0),
	}
	state.applier = func(e Event) {
		envelop := &EventEnvelope{
			Name:    ctx.namer.GetEventName(e),
			Payload: e,
		}

		Log.Debug("Applying %v to %T", envelop.Name, source)
		state.Events = append(state.Events, envelop)
	}
	source.SetEventApplier(state.applier)
	ctx.sources[source] = state
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
