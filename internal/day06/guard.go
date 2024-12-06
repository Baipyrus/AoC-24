package day06

import "slices"

type Guard struct {
	Position    Component
	Orientation int
}

func (g *Guard) Patrol(grid Grid, dirs []Component) int {
	var path []Component

	for true {
		velocity := dirs[g.Orientation]
		nextPos := AddComponents(g.Position, velocity)

		// Save position in path
		if !slices.Contains(path, g.Position) {
			path = append(path, g.Position)
		}

		// Rotate +90°
		if slices.Contains(grid.Obstacles, nextPos) {
			g.Orientation = (g.Orientation + 1) % 4
			continue
		}

		// Reached end
		if grid.OutOfBounds(nextPos) {
			break
		}

		// Move forward
		g.Position = nextPos
	}

	return len(path)
}
