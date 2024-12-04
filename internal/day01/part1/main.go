package day01_part1

import (
	"fmt"
	"slices"

	"github.com/Baipyrus/AoC-24/internal/day01"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  1 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	left, right := day01.ParseInput(input)

	slices.Sort(left)
	slices.Sort(right)

	var sum uint64
	for idx := range left {
		l := left[idx]
		r := right[idx]

		if l < r {
			sum += r - l
		} else {
			sum += l - r
		}
	}

	fmt.Printf("Total distance between lists: %d\n", sum)
}
