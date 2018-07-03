package wall

import (
	"sort"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// System moves the wall tiles to look like falling
type System struct {
	entities                entityList
	sortingNeeded           bool
	speed, traveled, bottom float32
}

// New adds the walls to the scene
func (s *System) New(w *ecs.World) {
	s.speed = 30
	s.bottom = 640
}

// Add adds an entity to the system
func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, wall *Component) {
	s.entities = append(s.entities, entity{basic, space, wall})
	s.sortingNeeded = true
}

// AddByInterface adds an entity that implements the Able interface to
// the system
func (s *System) AddByInterface(i ecs.Identifier) {
	o, _ := i.(Able)
	s.Add(o.GetBasicEntity(), o.GetSpaceComponent(), o.GetWallComponent())
}

// Remove removes an entity from the system. It doesn't do anything as there is
// no need to remove entites from the system.
func (s *System) Remove(basic ecs.BasicEntity) {}

// Update is called every frame.
func (s *System) Update(dt float32) {
	if s.sortingNeeded {
		sort.Sort(s.entities)
	}
	d := s.speed * dt
	s.traveled += d
	s.bottom -= d
	for _, e := range s.entities {
		e.SpaceComponent.Position.Subtract(engo.Point{X: 0, Y: d})
	}
	if s.traveled >= 32 {
		s.entities[0].SpaceComponent.Position.Y = s.bottom
		s.entities[1].SpaceComponent.Position.Y = s.bottom
		s.sortingNeeded = true
		s.traveled -= 32
		s.bottom += 32
	}
}
