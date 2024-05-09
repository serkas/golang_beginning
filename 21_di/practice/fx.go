package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	storage := NewMemStorage()

	service := &Service{storage: storage, done: make(chan struct{})}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go service.Run()

	<-ctx.Done()
	service.Stop()
}

type Storage interface {
	Store(value int64)
}

type Service struct {
	storage Storage
	done    chan struct{}
}

func (s *Service) Run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.done:
			fmt.Printf("shutdown")
			return
		case ts := <-ticker.C:
			// do some work
			val := ts.Unix() % 10
			s.storage.Store(val)
			fmt.Printf("val=%d\n", val)
		}
	}
}

func (s *Service) Stop() {
	close(s.done)
}

type MemStorage map[int64]struct{}

func NewMemStorage() MemStorage {
	return make(map[int64]struct{})
}

func (ms MemStorage) Store(val int64) {
	ms[val] = struct{}{}
}
