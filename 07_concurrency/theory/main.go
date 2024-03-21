package main

import (
	"fmt"
	"time"
)

func main() {
	//goroutines()
	channels()
}

func goroutines() {
	// let's start it in a goroutine
	go wordPrinter("hello 2")

	go wordPrinter("hello!")

	// goroutine function also can be defined as an inline function
	go func() {
		fmt.Println("single print")
	}()
	time.Sleep(time.Second)
}

func wordPrinter(word string) {
	for {
		fmt.Println(word)
		time.Sleep(time.Second)
	}
}
