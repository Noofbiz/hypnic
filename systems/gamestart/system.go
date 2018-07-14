package gamestart

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

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
		options.SaveOptions()
		engo.SetSceneByName("GameScene", true)
	}
}
