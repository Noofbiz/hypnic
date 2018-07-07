package messages

// ScoreType is the type of a GameoverMessage
const ScoreType string = "ScoreMessage"

// Score is sent when pointes are scored
type Score struct {
	Amount int
}

// Type implements the engo.Message interface
func (Score) Type() string {
	return ScoreType
}
