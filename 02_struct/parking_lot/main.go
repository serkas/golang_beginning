package main

import "fmt"

func main() {
	fmt.Println()

	// Stage 0. Create a Car struct
	// Add 1-3 struct fields with info about a car
	// Add a method to print info about the car

	c1 := Car{
		Color:        "blue",
		PlatesNumber: "AB1234PP",
	}

	fmt.Println(c1.GetInfo())

	// Stage 1.
	// Create a parking with 1 parking spot (a struct field) for a car.
	// Create a method to print the parking state.

	parking := Parking{
		Spot: &c1,
	}

	fmt.Printf("%v", parking)
	parking.PrintState()

	// Stage 2. Ensure Parking Spot is a pointer to a Car

	// Stage 3. Create a function to place a car in a parking spot.
	// Print some side info about the result of this action
	// What if the parking spot is already occupied?

	// Stage 4. We can place only cars there. What if are riding a bike?
	//
	// Create an interface Parkable and use it in the parking function
	//
}
