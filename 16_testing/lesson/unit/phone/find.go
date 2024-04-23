package phone

import (
	"regexp"
)

var phoneRegex = regexp.MustCompile(`\+?\d*\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}`)

func FindAll(input string) []string {
	return phoneRegex.FindAllString(input, -1)
}

func Dummy(input string) string {
	return ""
}
