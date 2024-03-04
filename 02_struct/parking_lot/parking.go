package main

import "fmt"

type Car struct {
	Color        string
	PlatesNumber string
}

func (c Car) GetInfo() string {
	return fmt.Sprintf("%s car with number %q", c.Color, c.PlatesNumber)
}

type Parking struct {
	Spot *Car
}

func (p Parking) PrintState() {
	fmt.Println("Parking with a single spot")
	if p.Spot != nil {
		fmt.Printf("The spot is occupied by: %s\n", p.Spot.GetInfo())
	}

}
