package menu

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type System struct {
	e entity
}

func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent) {
	s.e = entity{basic, space}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if engo.Input.Mouse.Action == engo.Press {
		x := engo.Input.Mouse.X + engo.ResizeXOffset/(2*engo.GetGlobalScale().X)
		y := engo.Input.Mouse.Y + engo.ResizeYOffset/(2*engo.GetGlobalScale().Y)
		if s.e.SpaceComponent.Contains(engo.Point{X: x, Y: y}) {
			engo.SetSceneByName("MainMenuScene", true)
		}
	}
}
