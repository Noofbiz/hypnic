package musicbx

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

type System struct {
	e       entity
	checked bool

	ctex, utex *common.Texture
}

func (s *System) New(w *ecs.World) {
	s.checked = options.TheOptions.BGM
	c, _ := common.LoadedSprite("checked.png")
	s.ctex = c
	u, _ := common.LoadedSprite("unchecked.png")
	s.utex = u
}

func (s *System) Add(basic *ecs.BasicEntity, mouse *common.MouseComponent,
	render *common.RenderComponent, audio *common.AudioComponent) {
	s.e = entity{basic, mouse, render, audio}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.e.Clicked {
		if options.TheOptions.SFX {
			s.e.AudioComponent.Player.Rewind()
			s.e.AudioComponent.Player.Play()
		}
		engo.Mailbox.Dispatch(messages.Music{
			Cb: true,
		})
		if s.checked {
			s.e.RenderComponent.Drawable = s.utex
			s.checked = false
			return
		}
		s.e.RenderComponent.Drawable = s.ctex
		s.checked = true
	}
}
