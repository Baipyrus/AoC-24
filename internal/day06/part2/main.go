package day06_part2

import (
	"fmt"
	"slices"

	"github.com/Baipyrus/AoC-24/internal/day06"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  6 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	// symbols[0] = guard (must be in clockwise order)
	// symbols[1] = obstacle
	// symbols[2] = path
	symbols := [][]rune{{'^', '>', 'v', '<'}, {'#'}, {'x'}}
	// Direction to walk in for every guard symbol, in order
	directions := []day06.Component{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0}}

	grid, guard := day06.ParseInput(input, symbols)

	var count uint64
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			pos := day06.Component{X: x, Y: y}

			// Skip if there's already an obstacle or the guard is there
			if slices.Contains(grid.Obstacles, pos) || guard.Position == pos {
				continue
			}

			// Clone the grid and add a new obstacle
			a := grid
			a.Obstacles = append(a.Obstacles, pos)

			// Clone the guard to avoid modifying its state
			b := guard
			if b.Loop(a, directions) {
				count++
			}
		}
	}

	fmt.Printf("Amount of possible obstructions: %d\n", count)
}
