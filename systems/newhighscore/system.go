package newhighscore

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type System struct {
	e       entity
	colors  []common.Text
	at      int
	elapsed float32
}

func (s *System) Add(basic *ecs.BasicEntity, render *common.RenderComponent) {
	red := &common.Font{
		URL:  "kenpixel_square.ttf",
		FG:   color.RGBA{255, 0, 0, 255},
		Size: 28,
	}
	red.CreatePreloaded()
	s.colors = append(s.colors, common.Text{
		Font: red,
		Text: "New High Score!",
	})
	orange := &common.Font{
		URL:  "kenpixel_square.ttf",
		FG:   color.RGBA{255, 165, 0, 255},
		Size: 28,
	}
	orange.CreatePreloaded()
	s.colors = append(s.colors, common.Text{
		Font: orange,
		Text: "New High Score!",
	})
	yellow := &common.Font{
		URL:  "kenpixel_square.ttf",
		FG:   color.RGBA{255, 255, 0, 255},
		Size: 28,
	}
	yellow.CreatePreloaded()
	s.colors = append(s.colors, common.Text{
		Font: yellow,
		Text: "New High Score!",
	})
	s.e = entity{basic, render}
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(dt float32) {
	s.elapsed += dt
	if s.elapsed > 0.08 {
		s.e.RenderComponent.Drawable = s.colors[s.at]
		s.at++
		if s.at >= len(s.colors) {
			s.at = 0
		}
		s.elapsed = 0
	}
}
