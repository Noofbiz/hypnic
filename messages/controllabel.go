package messages

// ControlLabelType is the type of a GameoverMessage
const ControlLabelType string = "ControlLabelMessage"

// ControlLabel is sent when a lose condition is met
type ControlLabel struct {
	Up bool
}

// Type implements the engo.Message interface
func (ControlLabel) Type() string {
	return ControlLabelType
}
