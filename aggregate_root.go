package es

// AggregateRootInterface is a wrapper containing all the convenients methods
// for an event sourced aggregate.
type AggregateRootInterface interface {
	AggregateRootID() string
	RaiseEvents() []Event
	Apply(e Event)
}

// AggregateRoot struct is a convenient wrapper with common methods that can be
// shared between all the aggregates.
type AggregateRoot struct {
	events   []Event
	playhead int64
}

// Apply an event to the aggregate
func (a *AggregateRoot) Apply(e Event) {
	a.events = append(a.events, e)
}

// RaiseEvents pulls in the applied events to be saved in the event store.
func (a *AggregateRoot) RaiseEvents() []Event {
	return a.events
}

// NewAggregateRoot instanciate an empty AggregateRoot
func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{
		events:   []Event{},
		playhead: 0,
	}
}
