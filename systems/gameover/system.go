package gameover

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/scenes/gameend"
)

// System handles the game over conditions
type System struct {
	gameover bool
	score    int
}

// New sets up the game over system
func (s *System) New(w *ecs.World) {
	engo.Mailbox.Listen(messages.GameOverType, func(m engo.Message) {
		_, ok := m.(messages.GameOver)
		if !ok {
			return
		}
		s.gameover = true
	})
	engo.Mailbox.Listen(messages.ScoreType, func(m engo.Message) {
		msg, ok := m.(messages.Score)
		if !ok {
			return
		}
		s.score += msg.Amount
	})
}

// Remove doesn't do anything as the system has no entities
func (s *System) Remove(e ecs.BasicEntity) {}

// Update is called once each frame. Checks if any lose conditions have been met.
func (s *System) Update(dt float32) {
	if s.gameover {
		engo.SetScene(&gameend.Scene{
			Score: s.score,
		}, true)
	}
}
