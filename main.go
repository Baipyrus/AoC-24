package main

import (
	"github.com/Baipyrus/AoC-24/internal/inputs"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

func main() {
	challenges := registry.Get()

	exec := inputs.GetChallenge(challenges)
	exec()
}
