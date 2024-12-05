package day05_part2

import (
	"fmt"
	"slices"

	"github.com/Baipyrus/AoC-24/internal/day05"
	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  5 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	ordering, updates := day05.ParseInput(input)

	var sum uint64
	for _, pages := range updates {
		invalid := validateUpdate(ordering, pages)
		sortUpdate(ordering, &pages)

		// Sum pagenumbers of "middle" entries in every valid update
		if invalid {
			middle := len(pages) / 2
			sum += pages[middle]
		}
	}

	fmt.Printf("Sum of valid page updates: %d\n", sum)
}

func sortUpdate(ordering [][]uint64, pages *[]uint64) {
	slices.SortFunc(*pages, func(lSelf uint64, rSelf uint64) int {
		for _, rule := range ordering {
			lOther := rule[0]
			rOther := rule[1]

			if lSelf == lOther && rSelf == rOther {
				return -1
			}
			if lSelf == rOther && rSelf == lOther {
				return 1
			}
		}
		return 0
	})
}

func validateUpdate(ordering [][]uint64, pages []uint64) bool {
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

	return invalid
}
