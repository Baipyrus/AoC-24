package day02_part1

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

	var count uint64
	fmt.Printf("Amount of safe reports: %d\n", count)
}
