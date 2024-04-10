package main

import "fmt"

func main() {
	a()

	fmt.Println("main finished")
}

func a() {
	defer func() {
		//if r := recover(); r != nil {
		//	fmt.Printf("recovered from \"%v\"\n", r)
		//}
	}()

	//panic("panic in a")

	b()
}

func b() {
	c()

	fmt.Println("some code in b")
}

func c() {
	defer func() {
		fmt.Println("deferred code in c")
	}()
	fmt.Println("some code in c")

	panic("panic in c")

	fmt.Println("some code in c 2")
}
