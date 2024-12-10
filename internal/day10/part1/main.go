package day10_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day 10 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var sum uint64
	fmt.Printf("Sum of trailhead scores on map: %d\n", sum)
}
