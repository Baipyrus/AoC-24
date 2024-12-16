package day10_part1

import (
	"fmt"
	"math"
	"slices"

	"github.com/Baipyrus/AoC-24/internal/day10"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day 10 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	grid := day10.ParseInput(input)
	sum := traverse(grid)

	fmt.Printf("Sum of trailhead scores on map: %d\n", sum)
}

func traverse(grid day10.Grid) (sum uint64) {
	heads := grid.GetFaces(0)

	for _, cell := range heads {
		tails := dijkstra(grid, cell)
		sum += uint64(len(tails))
	}

	return sum
}

func dijkstra(grid day10.Grid, start *day10.Component) (tails []*day10.Component) {
	// Initialize cell values
	for _, cell := range grid.Cells {
		cell.Distance = math.MaxUint64
		cell.Parent = nil
	}
	start.Distance = 0
	queue := append([]*day10.Component{}, grid.Cells...)

	// Scan neighbors of queued elements
	for len(queue) > 0 {
		// Get and remove closest cell
		current := day10.GetClosest(queue)
		day10.RemoveCell(&queue, current)

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
