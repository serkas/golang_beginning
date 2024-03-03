package main

import "fmt"

const (
	dinnerHour int = 4
	workHours  int = 9
)

type profession string

func (p profession) String() string {
	return fmt.Sprintf("profession: %s", string(p))
}

func main() {
	// // *** zero value
	// var (
	// 	hoursPast         int
	// 	hasDinnerHappened bool
	// )

	// hoursPast = 4

	// fmt.Printf("Zero values: %d %t\n", hoursPast, hasDinnerHappened)

	// // *** small simple program
	// for i := 0; i < workHours; i++ {
	// 	if hoursPast == dinnerHour {
	// 		hasDinnerHappened = true
	// 		fmt.Printf("\n!!! Dinner time !!!\n\n")
	// 	} else {
	// 		fmt.Printf("%d work hours past\n", hoursPast)
	// 		fmt.Printf("Has dinner happened: %t\n", hasDinnerHappened)
	// 	}
	// 	hoursPast++
	// }

	// *** numbers
	arg1 := 2.2
	arg2 := -19
	fmt.Println(calcSum(int(arg1), arg2))

	var arg3 uint64 = 200
	var arg4 uint32 = 9
	fmt.Println(calcSum(int(arg3), int(arg4)))

	// *** byte / rune
	var b1 byte = 3
	fmt.Printf("%d\n", b1)
	var b2 uint8 = 3
	fmt.Printf("%t\n", b1 == b2)

	R1 := 'Р'
	fmt.Printf("One symbol: %c, one number: %d\n", R1, R1)
	var R2 int32 = 1056
	fmt.Printf("%t\n", R1 == R2)

	// smile := '😊'
	// fmt.Printf("Emoji %c\n", smile)

	fmt.Printf("%T\n\n", "pharmacist")

	prof := profession("mechanic")
	fmt.Printf("%s\n", prof)
	fmt.Printf("%T\n", prof)

}

func calcSum(a, b int) int {
	return a + b
}

// formatting verbs: https://pkg.go.dev/fmt#hdr-Printing
// playground: https://go.dev/play/