package day03_part2

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  3 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	r := regexp.MustCompile(`mul\((-?\d{1,3}),(-?\d{1,3})\)|don't\(\)|do\(\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	var (
		sum int64
		use = true
	)

	for _, match := range matches {
		switch match[0] {
		case "don't()":
			use = false
		case "do()":
			use = true
		}

		if !use {
			continue
		}

		l, _ := strconv.ParseInt(match[1], 10, 64)
		r, _ := strconv.ParseInt(match[2], 10, 64)

		sum += l * r
	}

	fmt.Printf("Sum of multiplications: %d\n", sum)
}
