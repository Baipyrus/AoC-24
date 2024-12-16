package day10

import "strings"

func ParseInput(input string) (grid Grid) {
	lines := strings.Split(input, "\n")

	for j, line := range lines {
		if line == "" {
			continue
		}

		y := uint64(j)
		if grid.Height < y+1 {
			grid.Height = y + 1
		}

		for i, char := range line {
			x := uint64(i)
			if grid.Width < x+1 {
				grid.Width = x + 1
			}

			value := uint64(char - '0')
			cell := &Component{
				X: x, Y: y,
				Value: value}

			grid.Cells = append(grid.Cells, cell)
		}
	}

	return grid
}
