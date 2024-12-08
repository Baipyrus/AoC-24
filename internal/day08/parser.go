package day08

import "strings"

func ParseInput(input string) Grid {
	lines := strings.Split(input, "\n")
	var grid Grid

	for j, line := range lines {
		if line == "" {
			continue
		}

		for i, char := range line {
			if grid.Width < i+1 {
				grid.Width = i + 1
			}

			if char == '.' {
				continue
			}

			grid.Antennas = append(grid.Antennas, Antenna{
				Position:  Component{X: i, Y: j},
				Frequency: char,
			})
		}

		if grid.Height < j+1 {
			grid.Height = j + 1
		}
	}

	return grid
}
