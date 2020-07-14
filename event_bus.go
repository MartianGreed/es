package es

// EventBusInterface is an event bus interface
type EventBusInterface interface {
	Subscribe(e EventHandlerInterface)
	Dispatch(e Event)
}

// SimpleEventBus is an easy to use event bus
type SimpleEventBus struct {
	handlers []EventHandlerInterface
}

// Subscribe registers an event handler into the event bus
func (eb *SimpleEventBus) Subscribe(e EventHandlerInterface) {
	eb.handlers = append(eb.handlers, e)
}

// Dispatch sends an event into the event handlers
func (eb *SimpleEventBus) Dispatch(e Event) {
	for _, h := range eb.handlers {
		if !h.Supports(e) {
			continue
		}

		h.Handle(e)
	}
}
