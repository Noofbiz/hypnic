package gem

type Component struct{}

func (c *Component) GetGemComponent() *Component {
	return c
}
