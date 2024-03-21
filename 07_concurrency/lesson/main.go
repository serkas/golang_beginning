package main

import "time"

func main() {
	// Count visitors entering a shopping mall

	mainDoor := NewEventSource("MainDoor")
	parkingEntry := NewEventSource("ParkingEntry")

	events := make(chan Event)

	go mainDoor.GenerateEvents(3*time.Second, events)
	go parkingEntry.GenerateEvents(5*time.Second, events)

	// Counter
	counter := NewCounter()
	counter.Run(events)
}
