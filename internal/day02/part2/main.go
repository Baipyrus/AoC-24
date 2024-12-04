package day02_part2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  2 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var (
		count uint64
		log   = parseInput(input)
	)

	for _, values := range log {
		if checkReport(values) {
			count++
			continue
		}

		for idx := range values {
			modified := make([]uint64, 0, len(values)-1)
			modified = append(modified, values[:idx]...)
			modified = append(modified, values[idx+1:]...)

			if checkReport(modified) {
				count++
				break
			}
		}
	}

	fmt.Printf("Amount of safe reports: %d\n", count)
}

func checkReport(values []uint64) bool {
	var (
		direction int64
		previous  uint64
		safe      = true
	)

	for idx, v := range values {
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

	return safe
}

func parseInput(input string) [][]uint64 {
	var log [][]uint64
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		reports := strings.Split(line, " ")
		if len(reports) == 0 || reports[0] == "" {
			continue
		}

		var values []uint64
		for _, report := range reports {
			v, _ := strconv.ParseUint(report, 10, 64)
			values = append(values, v)
		}

		log = append(log, values)
	}

	return log
}
