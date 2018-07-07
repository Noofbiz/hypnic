package messages

// DamageType is the type for the damage message
const DamageType string = "DamageMessage"

// Damage is a message that records damage taken
type Damage struct {
	Amount float32
}

// Type implements the engo.Message interface
func (Damage) Type() string {
	return DamageType
}
