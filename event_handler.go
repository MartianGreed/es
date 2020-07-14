package es

// EventHandlerInterface is the interface to implement in order to react to
// the application dispatch events
type EventHandlerInterface interface {
	Supports(e Event) bool
	Handle(e Event)
}
