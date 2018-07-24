package creditroll

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Noofbiz/hypnic/options"
)

// System is the bullet system
type System struct {
	entities []entity
	speed    float32
	w        *ecs.World
}

func (s *System) New(w *ecs.World) {
	s.w = w
	s.speed = 25
}

// Add adds an entity to the System
func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, cred *Component) {
	s.entities = append(s.entities, entity{basic, space, cred})
}

// AddByInterface adds an entity that implements the bullet able interface
func (s *System) AddByInterface(i ecs.Identifier) {
	o, _ := i.(Able)
	s.Add(o.GetBasicEntity(), o.GetSpaceComponent(), o.GetCreditRollComponent())
}

// Remove removes an entity from the System
func (s *System) Remove(basic ecs.BasicEntity) {
	d := s.elementExists(basic)
	if d >= 0 {
		s.entities = append(s.entities[:d], s.entities[d+1:]...)
	}
}

// Update is called each frame
func (s *System) Update(dt float32) {
	for i := 0; i < len(s.entities); i++ {
		s.entities[i].SpaceComponent.Position.Subtract(engo.Point{
			X: 0,
			Y: s.speed * dt,
		})
		if s.entities[i].Position.Y < -20+options.YOffset {
			s.w.RemoveEntity(*s.entities[i].BasicEntity)
		}
	}
}

func (s *System) elementExists(basic ecs.BasicEntity) int {
	d := -1
	for i, e := range s.entities {
		if e.ID() == basic.ID() {
			d = i
			break
		}
	}
	return d
}
