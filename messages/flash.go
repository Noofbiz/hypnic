package messages

// FlashType is the type of a GameoverMessage
const FlashType string = "FlashMessage"

// Flash is sent when a lose condition is met
type Flash struct{}

// Type implements the engo.Message interface
func (Flash) Type() string {
	return FlashType
}
