package bullet

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/engo/math"

	"github.com/Noofbiz/hypnic/collisions"
	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

// System is the bullet system
type System struct {
	entities []entity

	w *ecs.World

	ss        *common.Spritesheet
	animation *common.Animation

	speed float32
}

// New is called when the system is added to the world
func (s *System) New(w *ecs.World) {
	s.w = w
	s.ss = common.NewSpritesheetFromFile("bullet.png", 15, 15)
	s.animation = &common.Animation{Name: "bullet", Frames: []int{0, 1, 2, 3, 2, 1}, Loop: true}
	s.speed = 2.5

	pewp, _ := common.LoadedPlayer("pew.wav")
	pewp.SetVolume(0.75 * options.TheOptions.SFXLevel)
	pew := sound{BasicEntity: ecs.NewBasic()}
	pew.AudioComponent.Player = pewp
	w.AddEntity(&pew)

	engo.Mailbox.Listen(messages.SpeedType, func(m engo.Message) {
		_, ok := m.(messages.Speed)
		if !ok {
			return
		}
		s.speed += 0.75
	})
	engo.Mailbox.Listen(messages.CreateBulletType, func(m engo.Message) {
		msg, ok := m.(messages.CreateBullet)
		if !ok {
			return
		}
		b := bullet{BasicEntity: ecs.NewBasic()}
		b.RenderComponent.Drawable = s.ss.Drawable(0)
		b.RenderComponent.Scale = engo.Point{X: 2, Y: 2}
		b.RenderComponent.SetZIndex(4)
		b.SpaceComponent.Position = msg.Position
		b.SpaceComponent.Width = b.RenderComponent.Drawable.Width() * b.RenderComponent.Scale.X
		b.SpaceComponent.Height = b.RenderComponent.Drawable.Height() * b.RenderComponent.Scale.Y
		b.Component.Angle = msg.Angle
		b.AnimationComponent = common.NewAnimationComponent(s.ss.Drawables(), 0.3)
		b.AnimationComponent.AddDefaultAnimation(s.animation)
		b.CollisionComponent.Group = collisions.Player
		w.AddEntity(&b)
		if options.TheOptions.SFX {
			pewp.Rewind()
			pewp.Play()
		}
	})
	engo.Mailbox.Listen("CollisionMessage", func(m engo.Message) {
		msg, ok := m.(common.CollisionMessage)
		if !ok {
			return
		}
		if s.elementExists(*msg.To.BasicEntity) >= 0 {
			engo.Mailbox.Dispatch(messages.Damage{
				Amount: -20,
			})
			w.RemoveEntity(*msg.To.BasicEntity)
		}
	})
}

// Add adds an entity to the System
func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, bullet *Component) {
	s.entities = append(s.entities, entity{basic, space, bullet})
}

// AddByInterface adds an entity that implements the bullet able interface
func (s *System) AddByInterface(i ecs.Identifier) {
	o, _ := i.(Able)
	s.Add(o.GetBasicEntity(), o.GetSpaceComponent(), o.GetBulletComponent())
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
		sin, cos := math.Sincos(s.entities[i].Angle)
		s.entities[i].Position.Subtract(engo.Point{
			X: s.speed * sin,
			Y: s.speed * cos,
		})
		if s.entities[i].Position.X < 31+options.XOffset {
			s.entities[i].Position.X = 32 + options.XOffset
			s.entities[i].Angle *= -1
		}
		if s.entities[i].Position.X > 255+options.XOffset {
			s.entities[i].Position.X = 254 + options.XOffset
			s.entities[i].Angle *= -1
		}
		if s.entities[i].Position.Y < options.YOffset {
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
