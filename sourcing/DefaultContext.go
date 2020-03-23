package sourcing

var (
	defaultContext = newDefaultContext()
)

func Create(source interface{}, version Version) EventSource {
	return defaultContext.Create(source, version)
}
