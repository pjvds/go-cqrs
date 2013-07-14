package sourcing

var (
	defaultContext = newDefaultContext()
)

func AttachNew(source EventSource) {
	defaultContext.AttachNew(source)
}

func AttachWithHistory(source EventSource, history []EventEnvelope) {
	defaultContext.AttachWithHistory(source, history)
}

func GetState(source EventSource) *SourceState {
	return defaultContext.GetState(source)
}

func Detach(source EventSource) {
	defaultContext.Detach(source)
}
