package control

// Component is the control component
type Component struct {
	XOff float32
}

// GetControlComponent retrieves the control component of the entity
func (c *Component) GetControlComponent() *Component {
	return c
}
