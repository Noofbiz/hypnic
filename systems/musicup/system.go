package musicup

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

type System struct {
	e entity
}

func (s *System) Add(basic *ecs.BasicEntity, mouse *common.MouseComponent, audio *common.AudioComponent) {
	s.e = entity{basic, mouse, audio}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.e.Clicked {
		if options.TheOptions.SFX {
			s.e.AudioComponent.Player.Rewind()
			s.e.AudioComponent.Player.Play()
		}
		engo.Mailbox.Dispatch(messages.Music{
			Amount: 0.1,
		})
		engo.Mailbox.Dispatch(messages.MusicLabel{
			Up: true,
		})
	}
}
