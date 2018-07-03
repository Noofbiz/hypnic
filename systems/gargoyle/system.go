package gargoyle

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// System is a system for the gargoyles
type System struct {
	w *ecs.World

	entities []entity

	speed, elapsed, wait float32
}

// New is called when the system is added to the World
func (s *System) New(w *ecs.World) {
	s.w = w
	s.speed = 30
	s.wait = 10
}

// Add adds an entity to the system
func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, gargoyle *Component) {
	s.entities = append(s.entities, entity{basic, space, gargoyle})
}

// AddByInterface adds an entity to the system that implements the gargoyle interface
func (s *System) AddByInterface(i ecs.Identifier) {
	o, _ := i.(Able)
	s.Add(o.GetBasicEntity(), o.GetSpaceComponent(), o.GetGargoyleComponent())
}

// Remove removes an entity from the system
func (s *System) Remove(basic ecs.BasicEntity) {
	d := -1
	for i, e := range s.entities {
		if e.ID() == basic.ID() {
			d = i
			break
		}
	}
	if d >= 0 {
		s.entities = append(s.entities[:d], s.entities[d+1:]...)
	}
}

// Update is called each frame
func (s *System) Update(dt float32) {
	s.elapsed += dt
	if s.elapsed > s.wait {
		r := rand.Intn(3)
		if r == 0 {
			s.addGargoyle()
		} else if r == 1 {
			s.wait -= 5 * rand.Float32()
			if s.wait < 5 {
				s.wait = 5
			}
		} else {
			s.addGargoyle()
			s.wait += 5 * rand.Float32()
			if s.wait > 15 {
				s.wait = 15
			}
		}
		s.elapsed = 0
	}
	for _, e := range s.entities {
		e.Position.Subtract(engo.Point{X: 0, Y: s.speed * dt})
		e.elapsed += dt
		if e.elapsed < e.charge && e.charges > 0 {
			//fire!
			e.elapsed = 0
			e.charges--
		}
	}
}

func (s *System) addGargoyle() {
	gs, _ := common.LoadedSprite("gargoyle.png")
	g := gargoyle{BasicEntity: ecs.NewBasic()}
	g.RenderComponent.Drawable = gs
	g.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: rand.Float32()*(engo.GameWidth()-64) + 32,
			Y: engo.GameHeight(),
		},
		Width:  g.RenderComponent.Drawable.Width(),
		Height: g.RenderComponent.Drawable.Height(),
	}
	g.Component.charge = rand.Float32()*5 + 1
	g.Component.charges = rand.Intn(5)
	s.w.AddEntity(&g)
}
