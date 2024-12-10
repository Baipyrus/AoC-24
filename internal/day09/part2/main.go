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

	filesystem := parseInput(input)
	compactBlocks(&filesystem)
	checksum := calcChecksum(filesystem)

	fmt.Printf("Checksum of compacted disk blocks: %d\n", checksum)
}

func calcChecksum(fs []File) (sum uint64) {
	// Sum products of index and file id
	var idx uint64 = 0

	for _, file := range fs {
		if file.Empty {
			idx += file.Size
			continue
		}

		for range file.Size {
			sum += uint64(idx) * file.Id
			idx++
		}
	}

	return sum
}

func compactBlocks(fs *[]File) {
	// Save slice for better readability
	filesystem := *fs

	for i := len(filesystem) - 1; i >= 0; i-- {
		// Look for (non-empty) file blocks
		current := filesystem[i]
		if current.Empty {
			continue
		}

		// Find left-most empty blocks
		j := slices.IndexFunc(filesystem[:i], func(f File) bool {
			return f.Empty && f.Size >= current.Size
		})
		if j == -1 {
			continue
		}

		// Get difference in size before change
		diff := filesystem[j].Size - current.Size

		// Move file blocks into empty slots
		filesystem[j].Id = current.Id
		filesystem[j].Size = current.Size
		filesystem[j].Empty = false

		// Mark the old blocks as empty
		filesystem[i].Empty = true

		if diff > 0 {
			sub := File{
				Empty: true,
				Size:  diff}
			// Insert the new empty block after the moved block
			filesystem = append(
				filesystem[:j+1],
				append(
					[]File{sub},
					filesystem[j+1:]...,
				)...)
		}
	}

	// Write the modified slice back to the pointer
	*fs = filesystem
}

func parseInput(input string) (fs []File) {
	var id uint64

	for idx, char := range input {
		// Invalid character
		if char < '0' || char > '9' {
			continue
		}

		amount := uint64(char - '0')
		empty := idx%2 == 1

		file := File{Empty: empty, Size: amount}

		if !empty {
			file.Id = id
			id++
		}

		fs = append(fs, file)
	}

	return fs
}

type File struct {
	Id    uint64
	Size  uint64
	Empty bool
}
