package day06_part1

import (
	"fmt"
	"slices"
	"strings"

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
	directions := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	var (
		obstacles, path [][]int
		position        []int
		width, height   int
		orientation     int
	)

	upper := strings.ToLower(input)
	lines := strings.Split(upper, "\n")
	for j, line := range lines {
		if line == "" {
			continue
		}

		for i, char := range line {
			if slices.Contains(symbols[0], char) {
				orientation = slices.Index(symbols[0], char)
				position = []int{i, j}
			} else if slices.Contains(symbols[1], char) {
				obstacles = append(obstacles, []int{i, j})
			}

			if width < j+1 {
				width = j + 1
			}
		}

		if height < j+1 {
			height = j + 1
		}
	}

	for true {
		velocity := directions[orientation]
		nextPos := day06.GuardWalk(position, velocity)

		// Save position in path
		if !day06.ContainsIntArr(path, position) {
			path = append(path, position)
		}

		// Rotate +90°
		if day06.ContainsIntArr(obstacles, nextPos) {
			orientation = (orientation + 1) % len(symbols[0])
			continue
		}

		// Reached end
		if day06.OutOfBounds(nextPos, width, height) {
			break
		}

		// Move forward
		position = nextPos
	}

	fmt.Printf("Distinct positions in mapped area: %d\n", len(path))
}
