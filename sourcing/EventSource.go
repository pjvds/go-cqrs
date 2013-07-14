package sourcing

type EventSource interface {
	SetEventApplier(applier EventHandler)
}
