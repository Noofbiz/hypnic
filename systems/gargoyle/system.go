package gargoyle

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"engo.io/engo/math"

	"github.com/Noofbiz/hypnic/messages"
)

// System is a system for the gargoyles
type System struct {
	w *ecs.World

	entities []entity

	playerposition engo.Point

	speed, elapsed, wait    float32
	max, min, increment     float32
	mincharges, randcharges int
}

// New is called when the system is added to the World
func (s *System) New(w *ecs.World) {
	s.w = w
	s.speed = 30
	s.wait = 7
	s.max = 10
	s.min = 3
	s.increment = 5
	s.mincharges = 1
	s.randcharges = 3
	engo.Mailbox.Listen(messages.SpeedType, func(m engo.Message) {
		_, ok := m.(messages.Speed)
		if !ok {
			return
		}
		s.max--
		s.min -= 0.15
		s.increment += 0.5
		s.mincharges++
		s.randcharges += 2
	})
	engo.Mailbox.Listen(messages.GetPlayerPositionType, func(m engo.Message) {
		msg, ok := m.(messages.GetPlayerPosition)
		if !ok {
			return
		}
		s.playerposition = msg.Position
	})
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
			s.wait -= s.increment * rand.Float32()
			if s.wait < s.min {
				s.wait = s.min
			}
		} else {
			s.addGargoyle()
			s.wait += s.increment * rand.Float32()
			if s.wait > s.max {
				s.wait = s.max
			}
		}
		s.elapsed = 0
	}
	for i := 0; i < len(s.entities); i++ {
		s.entities[i].Position.Subtract(engo.Point{X: 0, Y: s.speed * dt})
		s.entities[i].elapsed += dt
		if s.entities[i].elapsed > s.entities[i].charge && s.entities[i].charges > 0 {
			engo.Mailbox.Dispatch(messages.SendPlayerPosition{})
			pt := s.entities[i].SpaceComponent.Center()
			a := math.Atan((pt.X - s.playerposition.X) / (pt.Y - s.playerposition.Y))
			engo.Mailbox.Dispatch(messages.CreateBullet{
				Position: pt,
				Angle:    a,
			})
			s.entities[i].elapsed = 0
			s.entities[i].charges--
		}
		if s.entities[i].Position.Y < -66 {
			s.w.RemoveEntity(*s.entities[i].BasicEntity)
		}
	}
}

func (s *System) addGargoyle() {
	gs, _ := common.LoadedSprite("gargoyle.png")
	g := gargoyle{BasicEntity: ecs.NewBasic()}
	g.RenderComponent.Drawable = gs
	g.RenderComponent.SetZIndex(1)
	g.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: rand.Float32()*(engo.GameWidth()-64-g.RenderComponent.Drawable.Width()) + 32,
			Y: engo.GameHeight(),
		},
		Width:  g.RenderComponent.Drawable.Width(),
		Height: g.RenderComponent.Drawable.Height(),
	}
	g.Component.charge = rand.Float32() + 0.25
	g.Component.charges = rand.Intn(s.randcharges) + s.mincharges
	s.w.AddEntity(&g)
}
