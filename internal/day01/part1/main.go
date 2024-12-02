package day01_part1

import (
	"fmt"

	registry "github.com/Baipyrus/AoC-24/internal"
)

var name = "Day  1 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main() {
	fmt.Printf("Executing: %s\n", name)
}
