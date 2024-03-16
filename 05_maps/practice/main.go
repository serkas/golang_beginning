package main

import "fmt"

var m map[int]int

func main() {
	// Count number of repeats of each word in the sample text.
	// Use https://pkg.go.dev/strings#Split to split sting into separate words.

	sampleText := `There are some things that are so unforgivable that they make other things easily forgivable`

	wordCounts := countWords(sampleText)
	fmt.Println(wordCounts)

	// Index strings slice by words
	text := []string{
		"Hash map data structures use a hash function, which turns a key into an index within an underlying array.",
		"The hash function can be used to access an index when inserting a value or retrieving a value from a hash map.",
		"Hash map underlying data structure",
		"Hash maps are built on top of an underlying array data structure using an indexing system.",
		"Each index in the array can store one key-value pair.",
		"If the hash map is implemented using chaining for collision resolution, each index can store another data structure such as a linked list, which stores all values for multiple keys that hash to the same index.",
		"Each Hash Map key can be paired with only one value. However, different keys can be paired with the same value.",
	}

	fmt.Println(text)
}

func countWords(text string) map[string]int {
	return nil
}
