package flash

// FComponent is the component for the flash system
type FComponent struct{}

// GetFlashComponent returns the flash component of the entity
func (c *FComponent) GetFlashComponent() *FComponent {
	return c
}
