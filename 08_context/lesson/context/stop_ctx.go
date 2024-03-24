package main

import (
	"context"
	"fmt"
	"time"
)

func runStoppingWithContext() {
	// More reusable approach with context
	ctx, contextCancelFunc := context.WithCancel(context.Background())

	tasks := make(chan int)
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
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
		contextCancelFunc()
	}()

	for task := range tasks {
		fmt.Printf("task %d started\n", task)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("task %d done\n", task)
	}

	fmt.Println("done")
}
