package day01_part2

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/day01"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  1 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	left, right := day01.ParseInput(input)

	var sum uint64
	for _, l := range left {
		var count uint64
		for _, r := range right {
			if r == l {
				count++
			}
		}

		sum += count * l
	}

	fmt.Printf("Similarity score between lists: %d\n", sum)
}
