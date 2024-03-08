package main

import "testing"

func TestMaxOfThree(t *testing.T) {
	cases := []struct {
		input  [3]int
		result int
	}{
		{
			input:  [3]int{1, 4, 5},
			result: 5,
		},
		{
			input:  [3]int{6, 4, 5},
			result: 6,
		},
		{
			input:  [3]int{1, 8, 5},
			result: 8,
		},
		{
			input:  [3]int{5, 4, 5},
			result: 5,
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			result := maxOfThree(c.input[0], c.input[1], c.input[2])
			if result != c.result {
				t.Logf("expected %d, got %d", c.result, result)
				t.Fail()
			}
		})
	}
}
