package sourcing

var (
	defaultContext = newDefaultContext()
)

func CreateNew(source interface{}) EventSource {
	return defaultContext.CreateNew(source)
}

func CreateFromHistory(source interface{}, history []Event) EventSource {
	return defaultContext.CreateFromHistory(source, history)
}
