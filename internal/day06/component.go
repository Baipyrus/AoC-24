package day06

type Component struct {
	X int
	Y int
}

func (c *Component) Add(other Component) {
	c.X += other.X
	c.Y += other.Y
}

func AddComponents(a, b Component) Component {
	var n Component

	n.Add(a)
	n.Add(b)

	return n
}
