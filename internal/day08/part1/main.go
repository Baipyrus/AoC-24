package day08_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/day08"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  8 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	grid := day08.ParseInput(input)
	CalculateAntinodes(&grid)

	fmt.Printf("Unique locations within bounds: %d\n", len(grid.Antinodes))
}

func CalculateAntinodes(grid *day08.Grid) {
	groups := grid.GroupAntennas()

	for _, positions := range groups {
		for i, self := range positions {
			for j, other := range positions {
				if i == j {
					continue
				}

				// Only one direction needs to be
				// calculated because the other will
				// be implied on second iteration:
				// A - B + A
				diff := self.Sub(other)
				diff = diff.Add(self)
				// Other direction would be:
				// B - A + B

				bound := grid.OutOfBounds(diff)
				if !bound {
					grid.TryAddAntinode(diff)
				}
			}
		}
	}
}
