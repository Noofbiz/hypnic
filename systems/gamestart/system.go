package gamestart

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/scenes/game"
)

type System struct {
	e entity

	bgm, sfx     bool
	mlvl, sfxlvl float64
}

func (s *System) New(w *ecs.World) {
	s.bgm = true
	s.sfx = true
	s.mlvl = 1
	s.sfxlvl = 1

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

	engo.Mailbox.Listen(messages.SFXType, func(m engo.Message) {
		msg, ok := m.(messages.SFX)
		if !ok {
			return
		}
		if msg.Cb {
			s.sfx = !s.sfx
			return
		}
		s.sfxlvl += msg.Amount
		if s.sfxlvl >= 1 {
			s.sfxlvl = 0.99999
		}
		if s.sfxlvl <= 0 {
			s.sfxlvl = 0.00001
		}
	})
}

func (s *System) Add(basic *ecs.BasicEntity, mouse *common.MouseComponent) {
	s.e = entity{basic, mouse}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.e.Clicked {
		engo.SetScene(&game.Scene{
			BGM:    s.bgm,
			SFX:    s.sfx,
			MLvl:   s.mlvl,
			SFXLvl: s.sfxlvl,
		}, true)
	}
}
