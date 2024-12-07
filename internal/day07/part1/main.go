package day07_part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  7 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	lines := strings.Split(input, "\n")
	var equations []Equation

	for _, line := range lines {
		if line == "" {
			continue
		}

		members := strings.Split(line, " ")

		first := members[0]
		trim := first[:len(first)-1]
		result, _ := strconv.ParseInt(trim, 10, 64)

		var values []int64
		for i := 1; i < len(members); i++ {
			v, _ := strconv.ParseInt(members[i], 10, 64)
			values = append(values, v)
		}

		equations = append(equations, Equation{
			Result: result,
			Values: values})
	}

	var sum int64
	for _, eq := range equations {
		if eq.Validate() {
			sum += eq.Result
		}
	}

	fmt.Printf("Sum of valid calibration results: %d\n", sum)
}

type Equation struct {
	Result int64
	Values []int64
}

func (e *Equation) Validate() bool {
	var operators int64 = 2 // Add + Mult
	maximum := intPow(operators, int64(len(e.Values)-1))

	// State is a sequence of operators ("+, +, +" = 0 or "+, +, *" = 1 or ...)
	var state int64
	for state = range maximum {
		result := e.Values[0]

		// Go through every value
		for idx := int64(1); idx < int64(len(e.Values)); idx++ {
			curr := e.Values[idx]
			// Extract operator for current position from state
			op := (state / intPow(operators, idx-1)) % operators

			switch op {
			case 0:
				// Add to result
				result += curr
			case 1:
				// Multiply result
				result *= curr
			}
		}

		// If equation satisfied, return valid
		if result == e.Result {
			return true
		}
	}

	return false
}

func intPow(base, expo int64) int64 {
	if expo == 0 {
		return 1
	} else if expo == 1 {
		return base
	}

	result := base

	var iter int64
	for iter = 2; iter <= expo; iter++ {
		result *= base
	}
	return result
}
