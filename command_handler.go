package es

// CommandHandler interface
type CommandHandler interface {
	Supports(c Command) bool
	Handle(c Command)
}
