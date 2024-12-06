package day06_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/day06"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  6 - Part 1"

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
	distinct := guard.Patrol(grid, directions)

	fmt.Printf("Distinct positions in mapped area: %d\n", distinct)
}
