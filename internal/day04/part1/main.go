package day04_part1

import (
	"fmt"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  4 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	word := "XMAS"
	var count int
	fmt.Printf("Total appearances of '%s': %d\n", word, count)
}
