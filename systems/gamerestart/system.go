package gamerestart

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type System struct {
	e entity
}

func (s *System) Add(basic *ecs.BasicEntity, mouse *common.MouseComponent) {
	s.e = entity{basic, mouse}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.e.Clicked {
		engo.SetSceneByName("GameScene", true)
	}
}
