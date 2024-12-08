package day08

type Component struct {
	X int
	Y int
}

func (c *Component) Sub(o Component) Component {
	return Component{
		X: c.X - o.X,
		Y: c.Y - o.Y}
}

func (c *Component) Add(o Component) Component {
	return Component{
		X: c.X + o.X,
		Y: c.Y + o.Y}
}
