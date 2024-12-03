package day03_part2

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  3 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var sum int64
	fmt.Printf("Sum of multiplications: %d\n", sum)
}
