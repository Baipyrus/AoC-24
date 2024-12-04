package day04_part2

import (
	"fmt"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  4 - Part 2"

func init() {
	registry.Register(name, Main)
}

var dirs = [][]int{
	{1, 1},   // Bottom-Right
	{-1, 1},  // Bottom-Left
	{-1, -1}, // Top-Left
	{1, -1}}  // Top-Right

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	word := "MAS"
	upper := strings.ToUpper(input)
	lines := strings.Split(upper, "\n")

	var grid [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}

		chars := strings.Split(line, "")
		grid = append(grid, chars)
	}

	var count int
	for j := range grid {
		for i := range grid[j] {
			current := grid[j][i]
			if current != string(word[1]) {
				continue
			}

			combinations := [][]string{
				{getChar(grid, i, j, 0), current, getChar(grid, i, j, 2)},
				{getChar(grid, i, j, 1), current, getChar(grid, i, j, 3)},
				{getChar(grid, i, j, 2), current, getChar(grid, i, j, 0)},
				{getChar(grid, i, j, 3), current, getChar(grid, i, j, 1)},
			}

			var pattern int
			for _, i := range combinations {
				if strings.Join(i, "") == word {
					pattern++
				}
			}

			if pattern == 2 {
				count++
			}
		}
	}

	fmt.Printf("Total appearances of '%s': %d\n", word, count)
}

func getChar(grid [][]string, i, j, d int) string {
	x := i + dirs[d][0]
	y := j + dirs[d][1]

	yBound := y < 0 || y > len(grid)-1
	xBound := x < 0 || x > len(grid[j])-1

	if yBound || xBound {
		return ""
	}
	return grid[y][x]
}
