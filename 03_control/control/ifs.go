package main

import (
	"fmt"
	"time"
)

func ifs() {
	timestamp := time.Now().Unix() // Unix timestamp of a current moment
	fmt.Printf("Current timestamp: %d\n\n", timestamp)

	// Simple if
	// `{`,`}` are mandatory
	// `(`, `)` around condition are optional (and not recommended)
	if timestamp%2 == 0 {
		fmt.Printf("even second %d\n", timestamp)
	}

	//
	// If with an else
	if timestamp%2 == 0 {
		fmt.Printf("even second %d\n", timestamp)
	} else {
		fmt.Printf("odd second %d\n", timestamp)
	}

	// Usage of if
	fmt.Printf("max of three numbers %d\n", maxOfThree(3, 8, 2))

	fmt.Println()
}
