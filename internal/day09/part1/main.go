package day09_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-24/internal/registry"
)

var name = "Day  9 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	filesystem := parseInput(input)
	compactBlocks(&filesystem)
	checksum := calcChecksum(filesystem)

	fmt.Printf("Checksum of compacted disk blocks: %d\n", checksum)
}

func calcChecksum(fs []Block) (sum uint64) {
	// Sum products of index and file id
	for idx, file := range fs {
		if file.Empty {
			continue
		}

		sum += uint64(idx) * file.Id
	}

	return sum
}

func compactBlocks(fs *[]Block) {
	// Find the left most free space
	j := 0
	for idx := range *fs {
		if (*fs)[idx].Empty {
			j = idx
			break
		}
	}

	// Swap file blocks with free spaces, if any
	i := len(*fs) - 1
	for i > j {
		current := (*fs)[i]
		if !current.Empty {
			(*fs)[j] = current
			(*fs)[i].Empty = true

			// Find the next free space
			for j < len(*fs) && !(*fs)[j].Empty {
				j++
			}
		}
		i--
	}
}

func parseInput(input string) (fs []Block) {
	var id uint64

	for idx, char := range input {
		// Invalid character
		if char < '0' || char > '9' {
			continue
		}

		amount := uint64(char - '0')
		empty := idx%2 == 1

		// Expanding files into blocks early
		for range amount {
			file := Block{Empty: true}

			if !empty {
				file.Empty = false
				file.Id = id
			}

			fs = append(fs, file)
		}

		if !empty {
			id++
		}
	}

	return fs
}

type Block struct {
	Id    uint64
	Empty bool
}
