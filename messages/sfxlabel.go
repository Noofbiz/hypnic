package messages

// SFXLabelType is the type of a GameoverMessage
const SFXLabelType string = "SFXLabelMessage"

// SFXLabel is sent when a lose condition is met
type SFXLabel struct {
	Up bool
}

// Type implements the engo.Message interface
func (SFXLabel) Type() string {
	return MusicLabelType
}
