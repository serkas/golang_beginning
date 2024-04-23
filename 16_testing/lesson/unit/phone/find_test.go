package phone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllSingle(t *testing.T) {
	input := "number: +1(123)4567890"
	expected := []string{"+1(123)4567890"}

	result := FindAll(input)

	if len(result) != len(expected) {
		t.Fatalf("in: %s, expected: %s, got: %s", input, expected, result)
	}
	if result[0] != expected[0] {
		t.Errorf("in: %s, expected: %s, got: %s", input, expected, result)
	}
}

func TestFindAllTable(t *testing.T) {
	cases := []struct {
		name     string
		in       string
		expected []string
	}{
		{
			name:     "a",
			in:       "number: +1(123)4567890",
			expected: []string{"+1(123)4567890"},
		},
		{
			name:     "b",
			in:       "number: 123-456-7890 fdfsd",
			expected: []string{"123-456-7890"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := FindAll(c.in)

			if len(result) != len(c.expected) {
				t.Errorf("in: %s, expected: %s, got: %s", c.in, c.expected, result)
			}
			if result[0] != c.expected[0] {
				t.Errorf("in: %s, expected: %s, got: %s", c.in, c.expected, result)
			}
		})
	}
}

func TestFindAllTableAssert(t *testing.T) {
	cases := []struct {
		name     string
		in       string
		expected []string
	}{
		{
			name:     "with plus",
			in:       "number: +1(123)4567890",
			expected: []string{"+1(123)4567890"},
		},
		{
			name:     "with dashes",
			in:       "number: 123-456-7890 fdfsd",
			expected: []string{"123-456-7890"},
		},
		{
			name:     "two numbers (comma-separated)",
			in:       "number: 123-456-7890 +1(123)4567890 fdfsd",
			expected: []string{"123-456-7890", "+1(123)4567890"},
		},
		{
			name:     "two numbers (in text)",
			in:       "numbers +1(123)4567890 and 123-456-7890.",
			expected: []string{"+1(123)4567890", "123-456-7890"},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			result := FindAll(c.in)

			assert.Equal(t, c.expected, result)
		})
	}
}
