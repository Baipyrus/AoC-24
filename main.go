package main

import (
	"log"
	"os"

	"github.com/Baipyrus/AoC-24/internal/inputs"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

func main() {
	challenges := registry.Get()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	exec := inputs.GetChallenge(challenges)
	input := inputs.GetInput(cwd)

	exec(input)
}
