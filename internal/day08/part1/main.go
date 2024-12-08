package day08_part1

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  8 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	grid := parseInput(input)
	grid.CalculateAntinodes()

	fmt.Printf("Unique locations within bounds: %d\n", len(grid.Antinodes))
}

func parseInput(input string) Grid {
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

type Component struct {
	X int
	Y int
}

func (c *Component) Sub(o Component) Component {
	return Component{
		X: c.X - o.X,
		Y: c.Y - o.Y}
}

func (c *Component) Add(o Component) Component {
	return Component{
		X: c.X + o.X,
		Y: c.Y + o.Y}
}

type Antenna struct {
	Position  Component
	Frequency rune
}

type Grid struct {
	Antinodes []Component
	Antennas  []Antenna
	Width     int
	Height    int
}

func (g *Grid) OutOfBounds(c Component) bool {
	xBound := c.X < 0 || c.X > g.Width-1
	yBound := c.Y < 0 || c.Y > g.Height-1

	return xBound || yBound
}

func (g *Grid) GroupAntennas() map[rune][]Component {
	groups := make(map[rune][]Component)

	for _, antenna := range g.Antennas {
		groups[antenna.Frequency] = append(
			groups[antenna.Frequency],
			antenna.Position,
		)
	}

	return groups
}

func (g *Grid) CalculateAntinodes() {
	groups := g.GroupAntennas()

	for _, positions := range groups {
		for i, self := range positions {
			for j, other := range positions {
				if i == j {
					continue
				}

				// Only one direction needs to be
				// calculated because the other will
				// be implied on second iteration:
				// A - B + A
				diff := self.Sub(other)
				diff = diff.Add(self)
				// Other direction would be:
				// B - A + B

				bound := g.OutOfBounds(diff)
				exists := slices.Contains(g.Antinodes, diff)
				if !bound && !exists {
					g.Antinodes = append(g.Antinodes, diff)
				}
			}
		}
	}
}
