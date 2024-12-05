package day05

import (
	"strconv"
	"strings"
)

func ParseInput(input string) ([][]uint64, [][]uint64) {
	lines := strings.Split(input, "\n")
	var (
		ordering [][]uint64
		updates  [][]uint64
		toggle   bool
	)

	for _, line := range lines {
		// Blank-line to separate different inputs
		if line == "" {
			toggle = true
			continue
		}

		// Get rules
		if !toggle {
			pages := strings.Split(line, "|")
			l, _ := strconv.ParseUint(pages[0], 10, 64)
			r, _ := strconv.ParseUint(pages[1], 10, 64)

			ordering = append(ordering, []uint64{l, r})
			continue
		}

		// Get updates
		members := strings.Split(line, ",")
		var pages []uint64
		for _, page := range members {
			v, _ := strconv.ParseUint(page, 10, 64)
			pages = append(pages, v)
		}
		updates = append(updates, pages)
	}

	return ordering, updates
}
