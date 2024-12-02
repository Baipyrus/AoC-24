package inputs

import (
	"log"

	"github.com/Baipyrus/AoC-24/internal/registry"
	"github.com/ktr0731/go-fuzzyfinder"
)

type File struct {
	Path    string
	Content string
}

func GetChallenge(challenges []registry.Challenge) func() {
	idx, err := fuzzyfinder.Find(
		challenges,
		func(i int) string {
			return challenges[i].Name
		},
		fuzzyfinder.WithPromptString("Select Challenge: "))
	if err != nil {
		log.Fatal(err)
	}

	return challenges[idx].Exec
}
