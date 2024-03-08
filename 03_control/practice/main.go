package main

import "fmt"

func main() {
	// Practice 1. "Classic for + if"

	//Task: create a program to print out all numbers between 0 and N
	//that are multipliers of another number K (divided by K without a remainder).
	//Example: N = 20, K = 3; the result can be "3 6 9 12 15 18"
	//Hint: to get the remainder of division, use `%` operator, `xRemainder := x % K`

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// Practice 2. "range for + switch"
	// Count number of two characters in a string with a for loop in a single pass.
	//
	char1, char2 := 'a', 'c'
	text := "abcd abc cd"
	count1, count2 := countTwoCharacters(text, char1, char2)
	fmt.Printf("character `%c` occurred %d times, character `%c` occurred %d times\n", char1, count1, char2, count2)
}
