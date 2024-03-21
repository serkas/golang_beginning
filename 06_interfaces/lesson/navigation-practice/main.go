package main

import (
	"fmt"
	"navigation-practice/navigation"
	"navigation-practice/primitives"
	"navigation-practice/routing"
)

func main() {
	start := primitives.Point{
		X: 1,
		Y: 1,
	}

	finish := primitives.Point{
		X: 4,
		Y: 2,
	}

	nav := navigation.NewNavigator()
	nav.AddRoutingType("city", routing.CityRouter{})
	nav.AddRoutingType("land", routing.LandRouter{})

	fmt.Println("city navigation")
	instructions := nav.Navigate(start, finish, "city")
	for _, s := range instructions {
		fmt.Println(s)
	}

	fmt.Println()
	fmt.Println("land navigation")
	instructions2 := nav.Navigate(start, finish, "land")
	for _, s := range instructions2 {
		fmt.Println(s)
	}
}
