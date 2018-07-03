package messages

// GameOverType is the type of a GameoverMessage
const GameOverType string = "GameoverMessage"

// GameOver is sent when a lose condition is met
type GameOver struct{}

// Type implements the engo.Message interface
func (GameOver) Type() string {
	return GameOverType
}
