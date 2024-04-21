package main

import (
	"fmt"
	"strings"

	"proj/lessons/16_testing/lesson/unit/phone"
)

func main() {
	text := "text containing some phone numbers +1(123)4567890 and 123-456-7890. Addresses to 01001, 177 street"

	fmt.Println(strings.Join(phone.FindAll(text), ", "))
}
