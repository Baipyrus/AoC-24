package day12_part1

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day 12 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	regions := parseInput(input)
	var sum uint

	fmt.Printf("Amount of Regions: %d\n", len(regions))
	for _, r := range regions {
		area := uint(len(r.Plots))
		perimeter := r.Perimeter()
		price := area * perimeter
		sum += price

		fmt.Println()
		r.Println()
		fmt.Printf("Price of Region: %d * %d = %d\n", area, perimeter, price)
	}

	fmt.Printf("Total price of fencing all regions: %d\n", sum)
}

func parseInput(input string) (regions []*Region) {
	lines := strings.Split(input, "\n")
	counters := make(map[rune]uint)

	for j, line := range lines {
	loop:
		for i, char := range line {
			current := Component{X: i, Y: j}

			for _, r := range regions {
				// Plot not of the same type
				if char != r.Char {
					continue
				}
				if r.TryAdd(&current) {
					continue loop
				}
			}

			id := counters[char]
			next := Region{
				Id:    id,
				Char:  char,
				Plots: []*Component{&current}}
			regions = append(regions, &next)

			counters[char]++
		}
	}

	return regions
}

type Component struct {
	X int
	Y int
}

type Region struct {
	Id    uint
	Char  rune
	Plots []*Component
}

func (r *Region) Println() {
	fmt.Printf("#%d: %c -", r.Id, r.Char)
	for _, p := range r.Plots {
		fmt.Printf(" %d %d,", p.X, p.Y)
	}
	fmt.Println()
}

func (r *Region) Perimeter() (sum uint) {
	for _, a := range r.Plots {
		sum += 4 - r.Neighbors(a)
	}

	return sum
}

func (r *Region) TryAdd(c *Component) bool {
	// Plot already listed in Region
	if r.Contains(*c) {
		return false
	}

	neighboring := r.Neighbors(c) > 0

	// Add if neighbor exists
	if neighboring {
		r.Plots = append(r.Plots, c)
	}

	return neighboring
}

func (r *Region) Neighbors(c *Component) (count uint) {
	x, y := c.X, c.Y

	for _, n := range r.Plots {
		nx, ny := n.X, n.Y
		dx, dy := nx-x, ny-y

		//  0 = self
		//  2 = diagonal
		// >2 = invalid
		sum := dx*dx + dy*dy
		if sum == 1 {
			count++
		}
	}

	return count
}

func (r *Region) Contains(c Component) bool {
	return hasPlot(r.Plots, c)
}

func hasPlot(s []*Component, p Component) bool {
	return slices.ContainsFunc(
		s,
		func(c *Component) bool {
			return *c == p
		})
}
