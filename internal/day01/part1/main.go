package day01_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  1 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main() {
	fmt.Printf("Executing: %s\n", name)
}
