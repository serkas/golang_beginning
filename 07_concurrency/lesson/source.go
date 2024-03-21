package main

import (
	"fmt"
	"math/rand"
	"time"
)

type EventSource struct {
	name string
}

func NewEventSource(name string) *EventSource {
	return &EventSource{name: name}
}

func (s *EventSource) GenerateEvents(maxInterval time.Duration, events chan Event) {
	maxMicroseconds := maxInterval.Microseconds()

	var i int
	for {
		interval := time.Duration(rand.Int63n(maxMicroseconds)) * time.Microsecond
		time.Sleep(interval) // Using sleep here for simplicity. Should be replaced with time.Timer

		e := Event{ID: i, SourceName: s.name}
		i++
		fmt.Printf("%q generate event: %d\n", s.name, e.ID)
		events <- e
	}
}
