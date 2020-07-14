package es

// CommandBusInterface interface
type CommandBusInterface interface {
	Dispatch(c Command)
	Subscribe(c CommandHandler)
}

// SimpleCommandBus struct is a simple command bus for conveniance and fast
// design, you may want to change it for production use.
type SimpleCommandBus struct {
	handlers []CommandHandler
}

// Dispatch a command into the CommandBus if there are any listeners,
// they will be called concurrently
func (cb *SimpleCommandBus) Dispatch(c Command) {
	for _, handler := range cb.handlers {
		if !handler.Supports(c) {
			continue
		}
		go handler.Handle(c)
	}
}

// Subscribe adds a listener for any command, to make it work be sure to implement
// the Supports function
func (cb *SimpleCommandBus) Subscribe(c CommandHandler) {
	cb.handlers = append(cb.handlers, c)
}

// NewSimpleCommandBus Simple generator function
func NewSimpleCommandBus() *SimpleCommandBus {
	return &SimpleCommandBus{
		handlers: []CommandHandler{},
	}
}
