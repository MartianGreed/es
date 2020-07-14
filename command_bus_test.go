package es

import (
	"fmt"
	"testing"
)

type FakeCommand struct {
	Command
}

type FakeCommandHandler struct {
}

func (f FakeCommandHandler) Supports(c Command) bool {
	_, ok := c.(FakeCommand)
	return ok
}

func (f FakeCommandHandler) Handle(c Command) {
	fmt.Println("Hey I'm a fake command handler")
}

func TestSubscribeAddsToHandlers(t *testing.T) {
	cb := NewSimpleCommandBus()
	ch := FakeCommandHandler{}

	cb.Subscribe(ch)

	if 1 != len(cb.handlers) {
		t.Errorf("Failed to subscibe commandhandler")
	}
}

func TestDispatchCallsTheHandlerMethod(t *testing.T) {
	cb := NewSimpleCommandBus()
	ch := FakeCommandHandler{}

	cb.Subscribe(ch)

	cb.Dispatch(FakeCommand{})
}
