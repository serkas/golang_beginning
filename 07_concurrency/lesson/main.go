package main

import "time"

func main() {
	// Count visitors entering a shopping mall

	mainDoor := NewEventSource("MainDoor")
	//parkingEntry := NewEventSource("ParkingEntry")

	mainDoor.GenerateEvents(3 * time.Second)
	//parkingEntry.GenerateEvents(5 * time.Second)

}
