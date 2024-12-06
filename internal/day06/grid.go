package day06

type Grid struct {
	Obstacles []Component
	Width     int
	Height    int
}

func (g *Grid) OutOfBounds(c Component) bool {
	xBound := c.X < 0 || c.X > g.Width-1
	yBound := c.Y < 0 || c.Y > g.Height-1

	return xBound || yBound
}
