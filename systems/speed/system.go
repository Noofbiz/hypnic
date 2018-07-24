package speed

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/Noofbiz/hypnic/messages"
)

var timer = []float32{8, 8, 5, 4, 6}

type System struct {
	times   int
	elapsed float32
}

func (s *System) Remove(basic ecs.BasicEntity) {}

func (s *System) Update(dt float32) {
	if s.times < 5 {
		s.elapsed += dt
		if s.elapsed > timer[s.times] {
			engo.Mailbox.Dispatch(messages.Speed{})
			s.elapsed = 0
			s.times++
		}
	}
}
