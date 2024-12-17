package day12_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day 12 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)
}
