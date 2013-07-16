package sourcing

var (
	defaultContext = newDefaultContext()
)

func AttachNew(source interface{}) EventSource {
	return defaultContext.AttachNew(source)
}

func AttachFromHistory(source interface{}, history []EventEnvelope) EventSource {
	return defaultContext.AttachFromHistory(source, history)
}

func GetState(source interface{}) EventSource {
	return defaultContext.GetState(source)
}

func Detach(source interface{}) {
	defaultContext.Detach(source)
}
