package gargoyle

// Component is the gargoyle component
type Component struct {
	charge, elapsed float32
	charges         int
}

// GetGargoyleComponent returns the gargoyle component
func (c *Component) GetGargoyleComponent() *Component {
	return c
}
