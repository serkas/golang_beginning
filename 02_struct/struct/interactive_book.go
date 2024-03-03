package main

import "fmt"

const (
	BookInteractivityPuzzle       BookInteractivity = "puzzle"
	BookInteractivityOnlineQuizes BookInteractivity = "online_quizes"
)

type BookInteractivity string

type InteractiveBook struct {
	Book
	InteractivityType BookInteractivity
}

func (b InteractiveBook) ShortInfo() string {
	return fmt.Sprintf("%q (%d) by %s with %s interactivity", b.Title, b.Year, b.Author, b.InteractivityType)
}

func (b InteractiveBook) SetInteractivityType(interactivity BookInteractivity) {
	b.InteractivityType = interactivity // FIXME: broken assignment (pass-by-value receiver)
}

// func (b *InteractiveBook) SetInteractivityType(interactivity BookInteractivity) {
// 	b.InteractivityType = interactivity
// }
