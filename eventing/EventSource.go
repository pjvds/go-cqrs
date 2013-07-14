package eventing

type EventSource interface {
	SetEventApplier(applier EventHandler)
}
