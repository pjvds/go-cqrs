package sourcing

func newDefaultContext() *Context {
	return &Context{
		newRecorder: func(source interface{}) EventRecorder {
			return NewEventRecorder()
		},
		newRouterForSource: func(source interface{}) EventRouter {
			return NewReflectBasedRouter(source)
		},
	}
}

// The sourcing context.
type Context struct {
	newRecorder        func(source interface{}) EventRecorder
	newRouterForSource func(source interface{}) EventRouter
}

func (ctx *Context) create(source interface{}) EventSource {
	recorder := ctx.newRecorder(source)
	router := ctx.newRouterForSource(source)

	return newEventSource(router, recorder)
}

// Creates a new EventSource object that can be used to source events.
func (ctx *Context) CreateNew(source interface{}) EventSource {
	return ctx.create(source)
}

// Creates an existing EventSource object based on the state from the history
// are replays history so the specified source can update it's state.
func (ctx *Context) CreateFromHistory(source interface{}, history []Event) EventSource {
	eventSource := ctx.create(source)

	for _, event := range history {
		eventSource.Apply(event)
	}

	return eventSource
}
