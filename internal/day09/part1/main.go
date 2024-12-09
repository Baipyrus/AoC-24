package day09_part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  9 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	blocks := readDiskBlocks(input)
	compact := compactBlocks(blocks)
	sum := calculateChecksum(compact)

	fmt.Printf("Checksum of compacted disk blocks: %d\n", sum)
}

func calculateChecksum(input string) (sum int64) {
	// Sum products of index and file id
	for idx, char := range input {
		if char == '.' {
			continue
		}

		id := int(char - '0')
		prod := idx * id
		sum += int64(prod)
	}

	return sum
}

func compactBlocks(input string) string {
	blocks := []rune(input)

	// Find the left most free space
	j := 0
	for idx := range blocks {
		if blocks[idx] == '.' {
			j = idx
			break
		}
	}

	// Swap file blocks with free spaces, if any
	i := len(blocks) - 1
	for i > j {
		if blocks[i] != '.' {
			blocks[j], blocks[i] = blocks[i], '.'

			// Find the next free space
			for j < len(blocks) && blocks[j] != '.' {
				j++
			}
		}
		i--
	}

	return string(blocks)
}

func readDiskBlocks(input string) string {
	var (
		blocks strings.Builder
		id     uint64
	)

	for idx, char := range input {
		// Invalid character
		if char < '0' || char > '9' {
			continue
		}
		amount := int(char - '0')

		if idx%2 == 0 {
			// Append id to string x amount of times
			blocks.WriteString(strings.Repeat(
				strconv.FormatUint(id, 10),
				amount))

			id++
			continue
		}

		// Append free space, if any
		blocks.WriteString(strings.Repeat(".", amount))
	}

	return blocks.String()
}
