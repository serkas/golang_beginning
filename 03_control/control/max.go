package main

func maxOfThree(a, b, c int) int {
	abMax := a
	if a < b {
		abMax = b
	}

	if abMax > c {
		return abMax // it's a final result, so we can return directly here
	}

	return c // we don't need `else` subbranch, because we can reach here only if `abMax > c` is false
}
