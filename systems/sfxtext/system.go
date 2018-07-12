package sfxtext

import (
	"image/color"
	"math"
	"strconv"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

type System struct {
	e       entity
	text    int
	changed bool
	fnt     *common.Font
	texinit bool
	tex     common.Texture
}

func (s *System) New(w *ecs.World) {
	s.text = int(math.Round(options.TheOptions.SFXLevel * 10))

	engo.Mailbox.Listen(messages.MusicLabelType, func(m engo.Message) {
		msg, ok := m.(messages.SFXLabel)
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

	s.fnt = &common.Font{
		URL:  "Gaegu-Regular.ttf",
		FG:   color.White,
		Size: 32,
	}
	s.fnt.CreatePreloaded()
}

func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, render *common.RenderComponent) {
	s.e = entity{basic, space, render}
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
		if s.text == 2 {
			if !s.texinit {
				s.tex = s.fnt.Render("2")
				s.texinit = true
			}
			s.e.RenderComponent.SetShader(common.DefaultShader)
			s.e.RenderComponent.Drawable = s.tex
			s.e.SpaceComponent.Position.Y += 5
			s.changed = false
			return
		}
		t, ok := s.e.RenderComponent.Drawable.(common.Text)
		if !ok {
			s.e.RenderComponent.SetShader(common.TextShader)
			t = common.Text{
				Font: s.fnt,
			}
			s.e.SpaceComponent.Position.Y -= 5
		}
		t.Text = strconv.Itoa(s.text)
		s.e.RenderComponent.Drawable = t
		s.changed = false
	}
}
