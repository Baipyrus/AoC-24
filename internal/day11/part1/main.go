package day11_part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day 11 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	trimmed := strings.TrimSpace(input)
	splits := strings.Split(trimmed, " ")
	var numbers []uint64

	for _, n := range splits {
		parsed, _ := strconv.ParseUint(n, 10, 64)
		numbers = append(numbers, parsed)
	}

	for range 25 {
		var next []uint64

		for _, c := range numbers {
			n := applyRules(c)
			next = append(next, n...)
		}

		numbers = next
	}

	fmt.Printf("Amount of stones after blinking: %d\n", len(numbers))
}

func applyRules(input uint64) []uint64 {
	if input == 0 {
		return []uint64{1}
	}

	splits := trySplit(input)
	if len(splits) == 2 {
		return splits
	}

	var multiplier uint64 = 2024
	return []uint64{input * multiplier}
}

func trySplit(input uint64) []uint64 {
	// Return input number on uneven values
	amount := countDigits(input)
	if amount%2 == 1 {
		return []uint64{input}
	}

	// Create 'mask' for half the digits
	power := pow10(amount / 2)
	left := input / power
	right := input - left*power

	return []uint64{left, right}
}

func countDigits(input uint64) (count uint64) {
	// n = 0 not defined for log10(n) but len(fmt.Sprint(n)) = 1
	if input == 0 {
		return 1
	}

	// Keep dividing by 10 to simulate log10
	current := input
	for current > 0 {
		next := current / 10
		current = next
		count++
	}

	return count
}

func pow10(expo uint64) uint64 {
	var base uint64 = 10

	if expo == 0 {
		return 1
	} else if expo == 1 {
		return base
	}

	result := base

	var iter uint64
	for iter = 2; iter <= expo; iter++ {
		result *= base
	}
	return result
}
