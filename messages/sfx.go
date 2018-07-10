package messages

// SFXType is the type of a GameoverMessage
const SFXType string = "SFXMessage"

// SFX is sent when a lose condition is met
type SFX struct {
	Amount float64
	Cb     bool
}

// Type implements the engo.Message interface
func (SFX) Type() string {
	return SFXType
}
