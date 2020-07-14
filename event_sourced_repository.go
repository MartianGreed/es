package es

// EventSourcedRepositoryInterface is the interface you have to implement
// to interact with db aggregates or whatever
type EventSourcedRepositoryInterface interface {
	Load(id string)
	Save(ar AggregateRootInterface)
}
