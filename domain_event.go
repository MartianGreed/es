package es

import (
	"time"
)

// DomainEvent is a convenience wrapper arround event store events
type DomainEvent struct {
	ID         string
	Event      string
	Payload    string
	Metadata   string
	RecordedAt time.Time
	Playhead   int64
}

// DomainEventStream is a convenience wrapper for domain events
type DomainEventStream struct {
	Events []DomainEvent
}
