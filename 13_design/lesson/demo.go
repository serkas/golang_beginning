package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	runGame()
}

func runGame() {
	sender := NewQuestionsSender()
	// initiate scoring component

	// subscribe scoring component for questions

	// initiate players
	// provide players with channel for responses
	// register players in questions sender
	// start players

	// start scoring component

	sender.Start(context.Background())

	// implement graceful shutdown:
	//  - on signal, inform questions sender not to start the next round
	//  - wait for the end of the current round in scoring component
	//  - print the results
}

type Sender struct {
	questions []string
}

func (s *Sender) Start(ctx context.Context) {
	nextRoundTicker := time.NewTicker(time.Second * 10)

	for _, q := range s.questions {
		fmt.Println("send question", q)
		// send question to players here
		select {
		case <-ctx.Done():
			return
		case <-nextRoundTicker.C:
			continue
		}
	}
}

func NewQuestionsSender() *Sender {
	return &Sender{
		questions: []string{
			"q1",
			"q2",
		},
	}
}

type QuestionMessage struct {
	Question string
	Deadline time.Time
}
