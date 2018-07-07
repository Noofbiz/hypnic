package messages

// SpeedType is the type of a GameoverMessage
const SpeedType string = "SpeedMessage"

// Speed is sent when a lose condition is met
type Speed struct{}

// Type implements the engo.Message interface
func (Speed) Type() string {
	return SpeedType
}
