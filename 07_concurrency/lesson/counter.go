package main

import "time"

type Counter struct {
	countTotal    int
	countInterval int
	intervalStart time.Time
}

func CountEvents() {

}

func (c *Counter) addEvent() {
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
