package endcredits

import (
	"engo.io/ecs"
	"engo.io/engo"
)

type System struct{}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if engo.Input.Mouse.Action == engo.Press {
		engo.SetSceneByName("MainMenuScene", true)
	}
}
