package musictext

import (
	"strconv"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
)

type System struct {
	e       entity
	text    int
	changed bool
}

func (s *System) New(w *ecs.World) {
	s.text = 10

	engo.Mailbox.Listen(messages.MusicLabelType, func(m engo.Message) {
		msg, ok := m.(messages.MusicLabel)
		if !ok {
			return
		}
		s.changed = true
		if msg.Up {
			s.text++
		} else {
			s.text--
		}
	})
}

func (s *System) Add(basic *ecs.BasicEntity, render *common.RenderComponent) {
	s.e = entity{basic, render}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.changed {
		if s.text < 1 {
			s.text = 1
		}
		if s.text > 10 {
			s.text = 10
		}
		t := s.e.RenderComponent.Drawable.(common.Text)
		t.Text = strconv.Itoa(s.text)
		s.e.RenderComponent.Drawable = t
	}
}
