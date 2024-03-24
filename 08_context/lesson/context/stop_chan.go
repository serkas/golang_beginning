package main

import (
	"fmt"
	"time"
)

func runStoppingWithChannel() {
	// Example: stopping goroutine with a stop channel
	stop := make(chan struct{})

	tasks := make(chan int)
	go func() {
		for i := 0; ; i++ {
			select {
			case <-stop:
				fmt.Println("got stop signal. stop tasks generation")
				close(tasks)
				return
			default:
				tasks <- i
			}
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("sending stop signal")
		close(stop)
	}()

	for task := range tasks {
		fmt.Printf("task %d started\n", task)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("task %d done\n", task)
	}

	fmt.Println("done")
}
