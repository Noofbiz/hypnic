package potion

type Component struct {
	Charge int
}

func (c *Component) GetPotionComponent() *Component {
	return c
}
