package soundadjust

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

type System struct {
	e           entity
	current, to float64
	curcb, cb   bool
}

func (s *System) New(w *ecs.World) {
	s.current = options.TheOptions.BGMLevel
	s.to = options.TheOptions.BGMLevel
	s.curcb = options.TheOptions.BGM

	engo.Mailbox.Listen(messages.MusicType, func(m engo.Message) {
		msg, ok := m.(messages.Music)
		if !ok {
			return
		}
		if msg.Cb {
			s.cb = true
			return
		}
		s.to = s.current + msg.Amount
	})
}

func (s *System) Add(basic *ecs.BasicEntity, audio *common.AudioComponent) {
	s.e = entity{basic, audio}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.cb {
		s.curcb = !s.curcb
		options.TheOptions.SetBGM(s.curcb)
		if s.curcb {
			s.e.AudioComponent.Player.Play()
		} else {
			s.e.AudioComponent.Player.Pause()
		}
		s.cb = false
	}
	if s.current != s.to {
		if s.to <= 0 {
			s.to = 0.01
			s.current = 0.01
			options.TheOptions.SetBGMLevel(s.current)
			s.e.AudioComponent.Player.Pause()
			return
		}
		if s.to >= 1 {
			s.to = 0.99
			s.current = 0.99
		}
		s.e.AudioComponent.Player.SetVolume(s.to)
		s.current = s.to
		options.TheOptions.SetBGMLevel(s.current)
	}
}
