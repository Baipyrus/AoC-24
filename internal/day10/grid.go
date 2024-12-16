package day10

type Grid struct {
	Cells  []*Component
	Width  uint64
	Height uint64
}

func (g *Grid) GetCell(x, y int) *Component {
	return g.Cells[y*int(g.Width)+x]
}

func (g *Grid) GetNeighbors(cell *Component) (output []*Component) {
	x, y := int(cell.X), int(cell.Y)

	for _, dy := range []int{-1, 0, 1} {
		for _, dx := range []int{-1, 0, 1} {
			// Ignore self and diagonal
			sum := dx*dx + dy*dy
			if sum != 1 {
				continue
			}

			nx, ny := x+dx, y+dy

			// Detect out of bounds
			xBound := nx < 0 || nx >= int(g.Width)
			yBound := ny < 0 || ny >= int(g.Height)
			if xBound || yBound {
				continue
			}

			next := g.GetCell(nx, ny)
			output = append(output, next)
		}
	}

	return output
}

func (g *Grid) GetFaces(face uint64) (output []*Component) {
	for _, cell := range g.Cells {
		if cell.Value != face {
			continue
		}

		output = append(output, cell)
	}

	return output
}
