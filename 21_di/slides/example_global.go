package main

import "fmt"

var (
	B *ComponentB
)

func main() {
	B.do()
}

func init() { // "magic" function called once on program start (before main())
	B = &ComponentB{}
}

func (b *ComponentB) do() {
	fmt.Println("b doing things")
}
