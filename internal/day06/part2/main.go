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
	count := obstaclePlacements(grid, guard, directions)

	fmt.Printf("Amount of possible obstructions: %d\n", count)
}

func obstaclePlacements(grid day06.Grid, guard day06.Guard, directions []day06.Component) uint {
	var count uint
	for j := range grid.Height {
		for i := range grid.Width {
			pos := day06.Component{X: i, Y: j}

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
	return count
}
