package messages

// MusicType is the type of a GameoverMessage
const MusicType string = "MusicMessage"

// Music is sent when a lose condition is met
type Music struct {
	Amount float64
	Cb     bool
}

// Type implements the engo.Message interface
func (Music) Type() string {
	return MusicType
}
