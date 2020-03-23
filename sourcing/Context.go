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

// Creates a new EventSource object that can be used to source events.
func (ctx *Context) Create(source interface{}, version Version) EventSource {
	recorder := ctx.newRecorder(source)
	router := ctx.newRouterForSource(source)

	return newEventSource(version, router, recorder)
}
