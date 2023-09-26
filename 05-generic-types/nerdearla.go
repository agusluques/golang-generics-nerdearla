package main

import "fmt"

// type Workshop struct {
// 	title    string
// 	speakers []string
// 	repoURL  string
// }

type Talk struct {
	title   string
	speaker string
}

type Track struct {
	title string
	talks []Talk
}

func NewTrack(title string, talks []Talk) Track {
	return Track{title, talks}
}

func (t *Track) GetTalks() []Talk {
	return t.talks
}

func main() {
	talks := []Talk{
		{"Go for beginners", "Nahuel"},
		{"Javascript is awesome", "Nico"},
		{"How to be a ninja in Rust", "Juan"},
	}

	// workshops := []Workshop{
	// 	{"Generics in Go", []string{"Agus", "Nico"}, "https://github.com/agusluques/golang-generics-nerdearla"},
	// 	{"Building a CI/CD pipeline", []string{"Matias", "Jorge"}, "https://github.com/mati/ci-cd-pipeline"},
	// 	{"Working with Kubernetes", []string{"Ari", "Rafa"}, "https://github.com/rafinha/kubertenes-workshop"},
	// }

	track := NewTrack("Dev", talks)

	fmt.Printf("%+v\n", track.GetTalks())
}
