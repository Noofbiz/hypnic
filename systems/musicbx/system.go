package musicbx

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
)

type System struct {
	e       entity
	checked bool

	ctex, utex *common.Texture
}

func (s *System) New(w *ecs.World) {
	s.checked = true
	c, _ := common.LoadedSprite("checked.png")
	s.ctex = c
	u, _ := common.LoadedSprite("unchecked.png")
	s.utex = u
}

func (s *System) Add(basic *ecs.BasicEntity, mouse *common.MouseComponent, render *common.RenderComponent) {
	s.e = entity{basic, mouse, render}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.e.Clicked {
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
