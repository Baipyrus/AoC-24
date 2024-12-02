package main

import (
	"log"

	registry "github.com/Baipyrus/AoC-24/internal"
	"github.com/ktr0731/go-fuzzyfinder"
)

func main() {
	challenges := registry.Get()

	idx, err := fuzzyfinder.Find(
		challenges,
		func(i int) string {
			return challenges[i].Name
		},
		fuzzyfinder.WithPromptString("Select Challenge: "))
	if err != nil {
		log.Fatal(err)
	}

	challenges[idx].Exec()
}
