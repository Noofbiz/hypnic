package creditsbtn

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/scenes/credits"
)

type System struct {
	e entity

	bgm  bool
	mlvl float64
}

func (s *System) New(w *ecs.World) {
	s.bgm = true
	s.mlvl = 1

	engo.Mailbox.Listen(messages.MusicType, func(m engo.Message) {
		msg, ok := m.(messages.Music)
		if !ok {
			return
		}
		if msg.Cb {
			s.bgm = !s.bgm
			return
		}
		s.mlvl += msg.Amount
		if s.mlvl >= 1 {
			s.mlvl = 0.99999
		}
		if s.mlvl <= 0 {
			s.mlvl = 0.00001
		}
	})
}

func (s *System) Add(basic *ecs.BasicEntity, mouse *common.MouseComponent) {
	s.e = entity{basic, mouse}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.e.Clicked {
		engo.SetScene(&credits.Scene{
			BGM:      s.bgm,
			BGMLevel: s.mlvl,
		}, true)
	}
}
