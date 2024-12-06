package day06_part2

import (
	"fmt"
	"slices"
	"strings"

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
	directions := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	obst, p, face, w, h := parseInput(input, symbols)
	distinct := runGuardPatrol(obst, directions, p, face, w, h)

	fmt.Printf("Distinct positions in mapped area: %d\n", distinct)
}

func parseInput(input string, symbols [][]rune) ([][]int, []int, int, int, int) {
	var (
		obstacles     [][]int
		position      []int
		width, height int
		orientation   int
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

	return obstacles, position, orientation, width, height
}

func runGuardPatrol(obst, dirs [][]int, pos []int, face, w, h int) uint {
	var path [][]int

	for true {
		velocity := dirs[face]
		nextPos := guardWalk(pos, velocity)

		// Save position in path
		if !containsIntArr(path, pos) {
			path = append(path, pos)
		}

		// Rotate +90°
		if containsIntArr(obst, nextPos) {
			face = (face + 1) % 4
			continue
		}

		// Reached end
		if outOfBounds(nextPos, w, h) {
			break
		}

		// Move forward
		pos = nextPos
	}

	return uint(len(path))
}

func outOfBounds(pos []int, w, h int) bool {
	xBound := pos[0] < 0 || pos[0] > w-1
	yBound := pos[1] < 0 || pos[1] > h-1
	return xBound || yBound
}

func guardWalk(pos, dir []int) []int {
	return []int{pos[0] + dir[0], pos[1] + dir[1]}
}

func containsIntArr(array [][]int, value []int) bool {
	return slices.ContainsFunc(array, func(member []int) bool {
		// Compare elements
		for idx, cur := range member {
			if cur != value[idx] {
				return false
			}
		}
		return true
	})
}
