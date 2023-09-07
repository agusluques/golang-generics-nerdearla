package main

import "fmt"

type Workshop struct {
	title    string
	speakers []string
}

type Talk struct {
	title   string
	speaker string
}

type event interface {
	Workshop | Talk
}

type Track[T event] struct {
	title  string
	events []*T
}

func NewTrack[T event](title string, events []*T) *Track[T] {
	return &Track[T]{title, events}
}

// Gets first event and removes it from the list
func (t *Track[T]) NextEvent() (*T, bool) {
	if len(t.events) == 0 {
		return nil, false
	}
	next := t.events[0]
	t.events = t.events[1:]
	return next, true
}

func main() {
	talks := []*Talk{
		{"Go for beginners", "Nahuel"},
		{"Javascript is awesome", "Sebas"},
		{"How to be a ninja in Rust", "Juan"},
	}

	track := NewTrack[Talk]("Dev", talks)

	for {
		event, ok := track.NextEvent()
		if !ok {
			break
		}
		fmt.Printf("%+v\n", event)
	}

	workshops := []*Workshop{
		{"Generics in Go", []string{"Agus", "Nico"}},
		{"Building a CI/CD pipeline", []string{"Matias", "Jorge"}},
		{"Working with Kubernetes", []string{"Ari", "Rafa"}},
	}

	track2 := NewTrack[Workshop]("Wokshops", workshops)

	for {
		event, ok := track2.NextEvent()
		if !ok {
			break
		}
		fmt.Printf("%+v\n", event)
	}

}
