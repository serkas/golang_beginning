package main

import (
	"fmt"
	"time"
)

func loops() {
	// Classic for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("iteration: %d\n", i)
		fmt.Println()
	}

	// `i` is not accessible here, it was defined only inside the loop scope

	// Endless loop (we need to exit it with `break` or `return`)
	for {
		fmt.Println()
		time.Sleep(2 * time.Second)

		t := time.Now()

		if t.Unix()%10 == 0 {
			fmt.Printf("stopping clock at: %v\n", t)

			break // this will end the loop immediately
		}

		fmt.Printf("now is: %v\n", t)

		if t.Unix()%2 == 1 {
			continue // this will skip the rest of for-loop body and jump to the next iteration
		}

		fmt.Println("we are doing a lot of work here")
	}

	fmt.Println()

	var text string = "text 日本語 їю"
	var c rune = 'x'
	// iteration over string with range
	for _, char := range text {
		if c == char {
			fmt.Printf("character %c\n", char)
		}

		fmt.Printf("character %c\n", char)
	}

	fmt.Println()
}
