package main

import "fmt"

type Talk struct {
	title   string
	speaker string
}

type Track struct {
	title string
	talks []*Talk
}

func NewTrack(title string, talks []*Talk) *Track {
	return &Track{title, talks}
}

// Gets first talk and removes it from the list
func (t *Track) NextTalk() (*Talk, bool) {
	if len(t.talks) == 0 {
		return nil, false
	}
	next := t.talks[0]
	t.talks = t.talks[1:]
	return next, true
}

func main() {
	talks := []*Talk{
		{"Go for beginners", "Nahuel"},
		{"Javascript is awesome", "Nico"},
		{"How to be a ninja in Rust", "Juan"},
	}

	track := NewTrack("Dev", talks)

	for {
		talk, ok := track.NextTalk()
		if !ok {
			break
		}
		fmt.Printf("%+v\n", talk)
	}
}
