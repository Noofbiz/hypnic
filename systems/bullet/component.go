package bullet

// Component is the bullet component
type Component struct {
	Angle float32
}

// GetBulletComponent returns the bullet component
func (c *Component) GetBulletComponent() *Component {
	return c
}
