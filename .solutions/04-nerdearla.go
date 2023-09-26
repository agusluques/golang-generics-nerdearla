package main

import "fmt"

type Workshop struct {
	title    string
	speakers []string
	repoURL  string
}

type Talk struct {
	title   string
	speaker string
}

type Event interface {
	Talk | Workshop
}
type Track[E Event] struct {
	title  string
	events []E
}

func NewTrack[E Event](title string, events []E) Track[E] {
	return Track[E]{title, events}
}

func (t *Track[E]) GetEvents() []E {
	return t.events
}

func main() {
	talks := []Talk{
		{"Go for beginners", "Nahuel"},
		{"Javascript is awesome", "Nico"},
		{"How to be a ninja in Rust", "Juan"},
	}

	track := NewTrack("Dev", talks)

	fmt.Printf("%+v\n", track.GetEvents())

	workshops := []Workshop{
		{"Generics in Go", []string{"Agus", "Nico"}, "https://github.com/agusluques/golang-generics-nerdearla"},
		{"Building a CI/CD pipeline", []string{"Matias", "Jorge"}, "https://github.com/mati/ci-cd-pipeline"},
		{"Working with Kubernetes", []string{"Ari", "Rafa"}, "https://github.com/rafinha/kubertenes-workshop"},
	}

	trackOfWorkshops := NewTrack("Workshops", workshops)

	fmt.Printf("%+v\n", trackOfWorkshops.GetEvents())
}
