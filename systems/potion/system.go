package potion

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Noofbiz/hypnic/collisions"
	"github.com/Noofbiz/hypnic/messages"
)

type System struct {
	w                    *ecs.World
	entities             []entity
	speed, elapsed, wait float32
	max, min, increment  float32
	player               *common.Player
}

func (s *System) New(w *ecs.World) {
	s.w = w
	s.speed = 30
	s.wait = 15
	s.max = 20
	s.min = 5
	s.increment = 5

	s.player, _ = common.LoadedPlayer("potion.wav")
	p := sound{BasicEntity: ecs.NewBasic()}
	p.AudioComponent.Player = s.player
	w.AddEntity(&p)

	engo.Mailbox.Listen(messages.SpeedType, func(m engo.Message) {
		_, ok := m.(messages.Speed)
		if !ok {
			return
		}
		s.speed += 5
		s.max -= 1.75
		s.min -= 0.25
		s.increment += 0.5
	})

	engo.Mailbox.Listen("CollisionMessage", func(m engo.Message) {
		msg, ok := m.(common.CollisionMessage)
		if !ok {
			return
		}
		d := s.elementExists(*msg.To.BasicEntity)
		if d >= 0 {
			s.player.Rewind()
			s.player.Play()
			engo.Mailbox.Dispatch(messages.Damage{
				Amount: float32(s.entities[d].Component.Charge),
			})
			w.RemoveEntity(*msg.To.BasicEntity)
		}
	})
}

func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, potion *Component) {
	s.entities = append(s.entities, entity{basic, space, potion})
}

func (s *System) AddByInterface(i ecs.Identifier) {
	o, _ := i.(Able)
	s.Add(o.GetBasicEntity(), o.GetSpaceComponent(), o.GetPotionComponent())
}

func (s *System) Remove(basic ecs.BasicEntity) {
	d := s.elementExists(basic)
	if d >= 0 {
		s.entities = append(s.entities[:d], s.entities[d+1:]...)
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

func (s *System) Update(dt float32) {
	s.elapsed += dt
	if s.elapsed > s.wait {
		r := rand.Intn(3)
		if r == 0 {
			s.addPotion()
		} else if r == 1 {
			s.wait -= s.increment * rand.Float32()
			if s.wait < s.min {
				s.wait = s.min
			}
		} else {
			s.addPotion()
			s.wait += s.increment * rand.Float32()
			if s.wait > s.max {
				s.wait = s.max
			}
		}
		s.elapsed = 0
	}
	for i := 0; i < len(s.entities); i++ {
		s.entities[i].Position.Subtract(engo.Point{X: 0, Y: s.speed * dt})
		if s.entities[i].Position.Y < -32 {
			s.w.RemoveEntity(*s.entities[i].BasicEntity)
		}
	}
}

func (s *System) addPotion() {
	ps, _ := common.LoadedSprite("potion.png")
	p := potion{BasicEntity: ecs.NewBasic()}
	p.RenderComponent.Drawable = ps
	p.RenderComponent.SetZIndex(2)
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: rand.Float32()*(engo.GameWidth()-64-p.RenderComponent.Drawable.Width()) + 32,
			Y: engo.GameHeight(),
		},
		Width:  p.RenderComponent.Drawable.Width(),
		Height: p.RenderComponent.Drawable.Height(),
	}
	p.CollisionComponent.Group = collisions.Player
	p.Component.Charge = rand.Intn(15) + 5
	s.w.AddEntity(&p)
}
