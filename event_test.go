package es

import (
	"fmt"
	"testing"
)

type FakeEvent struct {
	test string
}

func (f *FakeEvent) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"test": f.test,
	}
}

func (f *FakeEvent) Deserialize(d map[string]interface{}) Event {
	return &FakeEvent{
		test: fmt.Sprint(d["test"]),
	}
}

func NewFakeEvent() *FakeEvent {
	return &FakeEvent{"ThisIsATest"}
}

func TestItCanBeInstanciatedFromAString(t *testing.T) {
	e := NewFakeEvent()

	sd := e.Serialize()
	de := e.Deserialize(sd)
	d, ok := de.(*FakeEvent)
	if !ok {
		t.Error("Failed converting data back to type")
	}
	if d.test != e.test {
		t.Error("Failed asserting deserialized data")
	}
}
