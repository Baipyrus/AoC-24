package day07

import (
	"strconv"
	"strings"
)

func ParseInput(input string) []Equation {
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

	return equations
}
