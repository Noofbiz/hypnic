package gameover

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"

	"github.com/Noofbiz/hypnic/messages"
)

// System handles the game over conditions
type System struct {
	gameover bool
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
}

// Remove doesn't do anything as the system has no entities
func (s *System) Remove(e ecs.BasicEntity) {}

// Update is called once each frame. Checks if any lose conditions have been met.
func (s *System) Update(dt float32) {
	if s.gameover {
		log.Println("game over. user wins. Nooooooo! Bob!")
		s.gameover = false
	}
}
