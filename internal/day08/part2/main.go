package day08_part2

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  8 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)
}