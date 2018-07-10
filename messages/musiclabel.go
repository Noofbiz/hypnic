package messages

// MusicLabelType is the type of a GameoverMessage
const MusicLabelType string = "MusicLabelMessage"

// MusicLabel is sent when a lose condition is met
type MusicLabel struct {
	Up bool
}

// Type implements the engo.Message interface
func (MusicLabel) Type() string {
	return MusicLabelType
}
