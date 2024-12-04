package day04_part1

import (
	"fmt"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  4 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	word := "XMAS"
	dirs := [][]int{
		{1, 0},   // Right
		{1, 1},   // Bottom-Right
		{0, 1},   // Bottom
		{-1, 1},  // Bottom-Left
		{-1, 0},  // Left
		{-1, -1}, // Top-Left
		{0, -1},  // Top
		{1, -1}}  // Top-Right

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
			if current != string(word[0]) {
				continue
			}

			for _, d := range dirs {
				idx := 1
				chars := []string{current}

				for k := range len(word) - 1 {
					x := i + d[0]*(k+1)
					y := j + d[1]*(k+1)

					yBound := y < 0 || y > len(grid)-1
					xBound := x < 0 || x > len(grid[j])-1
					if yBound || xBound {
						break
					}

					next := grid[y][x]
					if next != string(word[idx]) {
						break
					}

					chars = append(chars, next)
					idx++
				}

				if strings.Join(chars, "") == word {
					count++
				}
			}
		}
	}

	fmt.Printf("Total appearances of '%s': %d\n", word, count)
}
