package sourcing

var (
	defaultContext = newDefaultContext()
)

func CreateNew(source interface{}) EventSource {
	return defaultContext.CreateNew(source)
}

func CreateFromHistory(source interface{}, id EventSourceId, history []Event) EventSource {
	return defaultContext.CreateFromHistory(source, id, history)
}
