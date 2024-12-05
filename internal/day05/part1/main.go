package day05_part1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  5 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	ordering, updates := parseInput(input)

	/*
		Unused code to extend rule chains:

		Take rules 1|2 and 2|3. This implies a rule 1|3 because of
		the ordering 1|2|3. Resulting set should be 1|2, 2|3, 1|3.
		Only intended for usage on small input sets. AoC inputs
		already have pre-extended rule chains (duplicate lhs/rhs).
	*/
	// ordering = calculateChains(ordering)

	var sum uint64
	for _, pages := range updates {
		var invalid bool

	pageScan:
		for i, page := range pages {
			// Look for a rule containing page
			for _, rule := range ordering {
				l := rule[0]
				r := rule[1]
				if page != r {
					continue
				}

				// Invalid if any page exists on the right
				// of current while being at lhs of ruleset.
				for j := i + 1; j < len(pages); j++ {
					c := pages[j]
					if c != l {
						continue
					}

					invalid = true
					break pageScan
				}
			}
		}

		// Sum pagenumbers of "middle" entries in every valid update
		if !invalid {
			middle := len(pages) / 2
			sum += pages[middle]
		}
	}

	fmt.Printf("Sum of valid page updates: %d\n", sum)
}

func calculateChains(ordering [][]uint64) [][]uint64 {
	var rules [][]uint64

	for _, order := range ordering {
		if !containsUintArr(rules, order) {
			rules = append(rules, order)
		}

		lSelf := order[0]
		rSelf := order[1]

		for _, other := range rules {
			lOther := other[0]
			rOther := other[1]

			// 1|2 and 2|3 result in 1|3
			if lSelf == rOther {
				rules = append(rules, []uint64{lOther, rSelf})
			}
			// 2|3 and 1|2 resukt in 1|3
			if rSelf == lOther {
				rules = append(rules, []uint64{lSelf, rOther})
			}
		}
	}

	return rules
}

func containsUintArr(array [][]uint64, value []uint64) bool {
	return slices.ContainsFunc(array, func(member []uint64) bool {
		for idx, cur := range member {
			if idx > len(value)-1 || cur != value[idx] {
				return false
			}
		}
		// Requires 100% match of array
		return true
	})
}

func parseInput(input string) ([][]uint64, [][]uint64) {
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
