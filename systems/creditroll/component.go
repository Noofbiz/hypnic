package creditroll

// Component is the bullet component
type Component struct{}

// GetCreditRollComponent returns the bullet component
func (c *Component) GetCreditRollComponent() *Component {
	return c
}
