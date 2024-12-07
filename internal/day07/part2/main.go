package day07_part2

import (
	"fmt"
	"strconv"

	"github.com/Baipyrus/AoC-24/internal/day07"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  7 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	equations := day07.ParseInput(input)

	var sum int64
	for _, eq := range equations {
		if eq.Validate(3, handler) {
			sum += eq.Result
		}
	}

	fmt.Printf("Sum of valid calibration results: %d\n", sum)
}

func handler(op int64, result, current *int64) {
	switch op {
	case 0:
		// Add to result
		*result = *result + *current
	case 1:
		// Multiply result
		*result = *result * *current
	case 2:
		// Concat result
		str := fmt.Sprintf("%d%d", *result, *current)
		*result, _ = strconv.ParseInt(str, 10, 64)
	}
}
