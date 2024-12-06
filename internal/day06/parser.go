package day06

import (
	"slices"
	"strings"
)

func ParseInput(input string, symbols [][]rune) (Grid, Guard) {
	upper := strings.ToLower(input)
	lines := strings.Split(upper, "\n")

	var (
		grid  Grid
		guard Guard
	)

	for j, line := range lines {
		if line == "" {
			continue
		}

		for i, char := range line {
			current := Component{X: i, Y: j}

			if slices.Contains(symbols[0], char) {
				guard.Orientation = slices.Index(symbols[0], char)
				guard.Position = current
			} else if slices.Contains(symbols[1], char) {
				grid.Obstacles = append(grid.Obstacles, current)
			}

			if grid.Width < j+1 {
				grid.Width = j + 1
			}
		}

		if grid.Height < j+1 {
			grid.Height = j + 1
		}
	}

	return grid, guard
}
