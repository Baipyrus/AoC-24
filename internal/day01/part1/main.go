package day01_part1

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  1 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var left, right []uint64

	lines := regexp.MustCompile(`\r*\n`).Split(input, -1)
	for _, line := range lines {
		values := regexp.MustCompile(`\s+`).Split(line, -1)
		if len(values) != 2 {
			continue
		}

		l, _ := strconv.ParseUint(values[0], 10, 64)
		r, _ := strconv.ParseUint(values[1], 10, 64)

		left = append(left, l)
		right = append(right, r)
	}

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
