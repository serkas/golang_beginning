package main

func countTwoCharacters(text string, char1, char2 rune) (count1, count2 int) {
	for _, c := range text {
		switch {
		case c == char1:
			count1 += 1
		case c == char2:
			count2 += 1
		}
	}

	return
}
