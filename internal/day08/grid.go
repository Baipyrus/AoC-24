package day08

type Grid struct {
	Antinodes []Component
	Antennas  []Antenna
	Width     int
	Height    int
}

func (g *Grid) OutOfBounds(c Component) bool {
	xBound := c.X < 0 || c.X > g.Width-1
	yBound := c.Y < 0 || c.Y > g.Height-1

	return xBound || yBound
}

func (g *Grid) GroupAntennas() map[rune][]Component {
	groups := make(map[rune][]Component)

	for _, antenna := range g.Antennas {
		groups[antenna.Frequency] = append(
			groups[antenna.Frequency],
			antenna.Position,
		)
	}

	return groups
}
