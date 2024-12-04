package day02_part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  2 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var count uint64
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		reports := strings.Split(line, " ")
		if len(reports) == 0 || reports[0] == "" {
			continue
		}

		var (
			direction int64
			previous  uint64
			safe      = true
		)

		for idx, report := range reports {
			v, _ := strconv.ParseUint(report, 10, 64)
			if idx == 0 {
				previous = v
				continue
			}

			diff := int64(v - previous)
			up := diff > 0 && direction < 0
			down := diff < 0 && direction > 0
			tole := diff > 3 || diff < -3 || diff == 0
			if up || down || tole {
				safe = false
				break
			}

			direction = diff
			previous = v
		}

		if safe {
			count++
		}
	}

	fmt.Printf("Amount of safe reports: %d\n", count)
}
