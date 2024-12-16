package day10

import "slices"

type Component struct {
	X        uint64
	Y        uint64
	Value    uint64
	Distance uint64
	Parent   *Component
}

func RemoveCell(cells *[]*Component, cell *Component) {
	// Detect cell in slice
	idx := slices.Index(*cells, cell)
	if idx == -1 {
		return
	}

	// Remove out of order (overwrite at idx, pop last)
	(*cells)[idx] = (*cells)[len(*cells)-1]
	*cells = (*cells)[:len(*cells)-1]
}

func GetClosest(cells []*Component) *Component {
	current := cells[0]

	for _, cell := range cells {
		if cell.Distance >= current.Distance {
			continue
		}

		current = cell
	}

	return current
}
