package score

import (
	"image/color"
	"strconv"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

type System struct {
	e             entity
	fnt           *common.Font
	changed       bool
	elapsed, time float32
	score         int
}

func (s *System) New(w *ecs.World) {
	s.time = 2
	s.fnt = &common.Font{
		URL:  "kenpixel_square.ttf",
		FG:   color.White,
		Size: 16,
	}
	err := s.fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}
	t := text{BasicEntity: ecs.NewBasic()}
	t.RenderComponent.Drawable = common.Text{
		Font: s.fnt,
		Text: "Score: 0",
	}
	t.RenderComponent.SetZIndex(11)
	t.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 200 + options.XOffset,
			Y: 5 + options.YOffset,
		},
	}
	w.AddEntity(&t)
	s.e = entity{&t.BasicEntity, &t.RenderComponent}

	engo.Mailbox.Listen(messages.ScoreType, func(m engo.Message) {
		msg, ok := m.(messages.Score)
		if !ok {
			return
		}
		s.score += msg.Amount
		s.changed = true
	})
	engo.Mailbox.Listen(messages.SpeedType, func(m engo.Message) {
		_, ok := m.(messages.Speed)
		if !ok {
			return
		}
		s.time -= 0.2
	})
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(dt float32) {
	s.elapsed += dt
	if s.elapsed > s.time {
		engo.Mailbox.Dispatch(messages.Score{Amount: 10})
		s.elapsed = 0
	}
	if s.changed {
		t := s.e.RenderComponent.Drawable.(common.Text)
		t.Text = "Score: " + strconv.Itoa(s.score)
		s.e.RenderComponent.Drawable = t
		s.changed = false
	}
}
