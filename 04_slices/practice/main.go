package main

import "fmt"

func main() {
	// Initiate empty slice
	// Add 20-50 elements to the slice one by one
	// Print the slice length and capacity at each iteration
	var s []string
	for i := 0; i < 40; i++ {
		s = append(s, fmt.Sprintf("num %d", i))

		fmt.Printf("%d %d\n", len(s), cap(s))
	}

	// Average movie score
	movieScores := []int{6, 8, 3, 9, 10, 7, 8, 10}

	fmt.Println("Movie Scores")
	fmt.Println(movieScores)
	fmt.Printf("average score: %.2f", avg(movieScores))
}

func avg(scores []int) float64 {
	if len(scores) == 0 {
		return 0
	}

	var sum int
	for _, s := range scores {
		sum += s
	}

	return float64(sum) / float64(len(scores))
}
