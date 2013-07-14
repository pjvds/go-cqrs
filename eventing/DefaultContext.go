package eventing

var (
	defaultContext = newDefaultContext()
)

func Attach(source EventSource) {
	defaultContext.Attach(source)
}

func GetState(source EventSource) *SourceState {
	return defaultContext.GetState(source)
}

func Detach(source EventSource) {
	defaultContext.Detach(source)
}
