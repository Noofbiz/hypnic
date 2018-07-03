package wall

// Component is required for the system
type Component struct{}

// GetWallComponent impl;ements the Wallable interface
func (c *Component) GetWallComponent() *Component {
	return c
}
