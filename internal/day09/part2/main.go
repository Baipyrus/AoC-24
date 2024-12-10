package day09_part2

import (
	"fmt"
	"slices"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  9 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var checksum uint64
	fmt.Printf("Checksum of compacted disk blocks: %d\n", checksum)
}
