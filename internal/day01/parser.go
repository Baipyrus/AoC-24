package day01

import (
	"strconv"
	"strings"
)

func ParseInput(input string) ([]uint64, []uint64) {
	var left, right []uint64

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		values := strings.Split(line, "   ")
		if len(values) != 2 {
			continue
		}

		l, _ := strconv.ParseUint(values[0], 10, 64)
		r, _ := strconv.ParseUint(values[1], 10, 64)

		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}
