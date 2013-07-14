package eventing

// An methods that can handle an event. Most event handlers will receive
// only specific types and they should do the type assertion themself.
type EventHandler func(e Event)
