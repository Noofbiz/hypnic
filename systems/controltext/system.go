package controltext

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/Noofbiz/hypnic/controllers"
	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

type System struct {
	e        entity
	text     []string
	textures []common.Texture
	current  int
	changed  bool
	fnt      *common.Font
}

func (s *System) New(w *ecs.World) {
	s.fnt = &common.Font{
		URL:  "Gaegu-Regular.ttf",
		FG:   color.White,
		Size: 32,
	}
	s.fnt.CreatePreloaded()

	if controllers.HasKeyboard() {
		s.text = append(s.text, "Keyboard")
		s.textures = append(s.textures, s.fnt.Render("Keyboard"))
	}
	if controllers.HasMouse() {
		s.text = append(s.text, "Touch")
		s.textures = append(s.textures, s.fnt.Render("Touch"))
	}
	if controllers.HasAccelerometer() {
		s.text = append(s.text, "Acceler")
		s.textures = append(s.textures, s.fnt.Render("Acceler"))
	}
	for i := 0; i < len(s.text); i++ {
		if options.TheOptions.Controls == s.text[i] {
			s.current = i
			break
		}
	}

	engo.Mailbox.Listen(messages.ControlLabelType, func(m engo.Message) {
		msg, ok := m.(messages.ControlLabel)
		if !ok {
			return
		}
		s.changed = true
		if msg.Up {
			s.current++
		} else {
			s.current--
		}
	})
}

func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, render *common.RenderComponent) {
	s.e = entity{basic, space, render}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(float32) {
	if s.changed {
		if s.current >= len(s.text) {
			s.current = len(s.text) - 1
		}
		if s.current < 0 {
			s.current = 0
		}
		s.e.RenderComponent.Drawable = s.textures[s.current]
		options.TheOptions.SetControls(s.text[s.current])
		s.changed = false
	}
}
