package main

import "fmt"

func maps() {
	// declare a map (not really usable yet)
	var myMap map[string]int
	fmt.Println(myMap)

	// create a map
	myMap = make(map[string]int)
	myMap["a"] = 1 // store a key-value pair
	fmt.Println(myMap)

	// create a map with literal
	planetsRadiusKm := map[string]int{
		"Mercury": 2440,
		"Venus":   6052,
		"Earth":   6371,
		"Mars":    3390,
		"Jupiter": 69911,
		"Saturn":  58232,
		"Uranus":  25362,
		"Neptune": 24622,
		"Pluto":   1188,
	}
	fmt.Println(planetsRadiusKm)

	// get element from map
	planet := "Mars"
	r := planetsRadiusKm[planet] // blind get
	fmt.Printf("%s radius is %dkm\n", planet, r)

	r, ok := planetsRadiusKm["Saturn"]
	fmt.Println(ok, r)

	r, ok = planetsRadiusKm["Moon"]
	fmt.Println(ok, r)

	// Update element (same as assignment)
	myMap["a"] = 12
	fmt.Println(myMap)

	// Remove element from map
	delete(planetsRadiusKm, "Pluto")
	fmt.Println(planetsRadiusKm)
}
