package main

import (
	"fmt"
	"time"
)

type Counter struct {
	countTotal    int
	countInterval int
	intervalStart time.Time
}

func NewCounter() *Counter {
	return &Counter{
		intervalStart: time.Now(),
	}
}

func (c *Counter) Run(events chan Event) {
	//for e := range events {
	//	c.addEvent(e)
	//	fmt.Println("counter stored an event")
	//}

	displayTicker := time.NewTicker(10 * time.Second)
	for {
		select {
		case e := <-events:
			c.addEvent(e)
		case <-displayTicker.C:
			c.displayStats()
			c.resetInterval()
		}
	}
}

func (c *Counter) addEvent(e Event) {
	c.countTotal += 1
	c.countInterval += 1
}

func (c *Counter) getIntervalStats() (count int, duration time.Duration) {
	return c.countInterval, time.Since(c.intervalStart)
}

func (c *Counter) resetInterval() {
	c.countInterval = 0
	c.intervalStart = time.Now()
}

func (c *Counter) displayStats() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("\t\t Mall visitors stats")
	fmt.Printf("\t Total visitors: %d\n", c.countTotal)
	var rate float64
	intervalCount, duration := c.getIntervalStats()
	if duration > 0 {
		rate = float64(intervalCount) / duration.Seconds()
	}
	fmt.Printf("\t Current rate: %.3f visitors per second\n", rate)
	fmt.Println("---------------------------------------------------")
}
