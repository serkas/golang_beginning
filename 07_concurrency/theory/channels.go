package main

import (
	"fmt"
)

func channels() {
	//ch := make(chan string)
	//
	//go func() {
	//	ch <- "s" // sending to a channel
	//	// try closing the channel
	//	close(ch)
	//}()
	//// reading from channel
	//a := <-ch
	//fmt.Printf("got: %q\n", a)
	//
	//// read more
	//b := <-ch
	//fmt.Printf("got: %q\n", b)
	//
	//// double assignment // try closing the channel
	//c, open := <-ch
	//fmt.Println(c, open)

	// read from channel in a loop
	done := make(chan error)
	ch2 := make(chan int)
	go func() {
		for i := range ch2 {
			fmt.Println(i)
		}
		fmt.Println("printing loop done")
		close(done)
	}()

	for i := 0; i < 100; i++ {
		ch2 <- i
	}

	close(ch2)

	<-done
}
