package day06

import "slices"

type Guard struct {
	Position    Component
	Orientation int
}

func (g *Guard) Patrol(grid Grid, dirs []Component) int {
	var path []Component

	for {
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
			return len(path)
		}

		// Move forward
		g.Position = nextPos
	}
}

type Memory struct {
	Position    Component
	Orientation int
}

func (g *Guard) Loop(grid Grid, dirs []Component) bool {
	path := make(map[Memory]bool)

	for {
		velocity := dirs[g.Orientation]
		nextPos := AddComponents(g.Position, velocity)

		// Save the current state
		memory := Memory{
			Position:    g.Position,
			Orientation: g.Orientation,
		}

		// Loop detected
		if path[memory] {
			return true
		}
		path[memory] = true

		// Rotate +90°
		if slices.Contains(grid.Obstacles, nextPos) {
			g.Orientation = (g.Orientation + 1) % 4
			continue
		}

		// Reached end
		if grid.OutOfBounds(nextPos) {
			return false
		}

		// Move forward
		g.Position = nextPos
	}
}
