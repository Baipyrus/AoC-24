package day08_part2

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/day08"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  8 - Part 2"

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

				grid.TryAddAntinode(self)
				diff := self.Sub(other)
				previous := self

				for {
					next := diff.Add(previous)

					bound := grid.OutOfBounds(next)
					if bound {
						break
					}

					grid.TryAddAntinode(next)
					previous = next
				}
			}
		}
	}
}
