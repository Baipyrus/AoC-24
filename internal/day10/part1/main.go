package day10_part1

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day 10 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	grid := ParseInput(input)
	sum := traverse(grid)

	fmt.Printf("Sum of trailhead scores on map: %d\n", sum)
}

func traverse(grid Grid) (sum uint64) {
	heads := grid.GetFaces(0)

	for _, cell := range heads {
		tails := dijkstra(grid, cell)
		sum += uint64(len(tails))
	}

	return sum
}

func dijkstra(grid Grid, start *Component) (tails []*Component) {
	// Initialize cell values
	for _, cell := range grid.Cells {
		cell.Distance = math.MaxUint64
		cell.Parent = nil
	}
	start.Distance = 0
	queue := append([]*Component{}, grid.Cells...)

	// Scan neighbors of queued elements
	for len(queue) > 0 {
		// Get and remove closest cell
		current := GetClosest(queue)
		RemoveCell(&queue, current)

		for _, next := range grid.GetNeighbors(current) {
			// Ignore already scanned
			scope := !slices.Contains(queue, next)
			// Ignore established ends
			exists := slices.Contains(tails, next)
			// Must increment face-value
			value := next.Value != current.Value+1
			if scope || exists || value {
				continue
			}

			// Decrease distance if shorter path
			dist := current.Distance + 1
			if dist < next.Distance {
				next.Distance = dist
				next.Parent = current
			}

			// Detect end-point
			if next.Value == 9 && next.Distance == 9 {
				tails = append(tails, next)
			}
		}
	}

	return tails
}

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

type Component struct {
	X        uint64
	Y        uint64
	Value    uint64
	Distance uint64
	Parent   *Component
}

func RemoveCell(cells *[]*Component, cell *Component) {
	// Detect cell in slice
	idx := slices.Index(*cells, cell)
	if idx == -1 {
		return
	}

	// Remove out of order (overwrite at idx, pop last)
	(*cells)[idx] = (*cells)[len(*cells)-1]
	*cells = (*cells)[:len(*cells)-1]
}

func GetClosest(cells []*Component) *Component {
	current := cells[0]

	for _, cell := range cells {
		if cell.Distance >= current.Distance {
			continue
		}

		current = cell
	}

	return current
}

type Grid struct {
	Cells  []*Component
	Width  uint64
	Height uint64
}

func (g *Grid) GetCell(x, y int) *Component {
	return g.Cells[y*int(g.Width)+x]
}

func (g *Grid) GetNeighbors(cell *Component) (output []*Component) {
	x, y := int(cell.X), int(cell.Y)

	for _, dy := range []int{-1, 0, 1} {
		for _, dx := range []int{-1, 0, 1} {
			// Ignore self and diagonal
			sum := dx*dx + dy*dy
			if sum != 1 {
				continue
			}

			nx, ny := x+dx, y+dy

			// Detect out of bounds
			xBound := nx < 0 || nx >= int(g.Width)
			yBound := ny < 0 || ny >= int(g.Height)
			if xBound || yBound {
				continue
			}

			next := g.GetCell(nx, ny)
			output = append(output, next)
		}
	}

	return output
}

func (g *Grid) GetFaces(face uint64) (output []*Component) {
	for _, cell := range g.Cells {
		if cell.Value != face {
			continue
		}

		output = append(output, cell)
	}

	return output
}
