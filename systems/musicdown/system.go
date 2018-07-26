package musicdown

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

func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, audio *common.AudioComponent) {
	s.e = entity{basic, space, audio}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if engo.Input.Mouse.Action == engo.Press {
		x := engo.Input.Mouse.X + engo.ResizeXOffset/(2*engo.GetGlobalScale().X)
		y := engo.Input.Mouse.Y + engo.ResizeYOffset/(2*engo.GetGlobalScale().Y)
		if s.e.SpaceComponent.Contains(engo.Point{X: x, Y: y}) {
			if options.TheOptions.SFX {
				s.e.AudioComponent.Player.Rewind()
				s.e.AudioComponent.Player.Play()
			}
			engo.Mailbox.Dispatch(messages.Music{
				Amount: -0.1,
			})
			engo.Mailbox.Dispatch(messages.MusicLabel{
				Up: false,
			})
		}
	}
}
