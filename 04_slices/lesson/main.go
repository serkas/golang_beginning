package main

import "fmt"

func main() {
	// Array - fixed (not resizable) list of elements
	a := [5]int32{5, 4, 3, 2, 1}
	fmt.Println(a)
	for i := range a {
		fmt.Printf("index=%d, value=%d, address=%v\n", i, a[i], &a[i])
	}

	// Slice - dynamic list (uses arrays for storing actual values)
	s1 := make([]int, 0, 5) // create an empty slice with length 0 and capacity 10
	fmt.Println("s1", cap(s1))
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4, 5} // create a slice with literal notation
	fmt.Println("s2", cap(s1))
	fmt.Println(s2)

	s3 := s1[1:3] // create a slice from another slice (it has some side effects)
	fmt.Println("s3", cap(s3))
	fmt.Println(s3)

	// Slice append (appending values to the empty slice s1
	for i := 0; i < 10; i++ {
		s1 = append(s1, i*10)
	}
	fmt.Println("elements added to s1")
	fmt.Println(s1)

	// Usage of the same array by two slices

	fmt.Println("Check what is in s3")
	fmt.Println(s3)
	fmt.Println("s3 uses the same underlying array as s1 (at least so far)")

	fmt.Println("Let's try to add some elements to s3") //
	for i := 3; i < 10; i++ {
		s3 = append(s3, -i*10)
	}
	fmt.Println(s3)

	fmt.Println("Let's check s1")
	fmt.Println(s1)
	fmt.Println("s1 was resized (data copied) and is not affected by appends to s3")
	fmt.Println()

	// copy
	source := []int{1, 2, 3, 4, 5}
	var destination = make([]int, 3)
	n := copy(destination, source)
	fmt.Printf("%d elements copied: %v", n, destination)
}
