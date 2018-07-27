package life

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

// System is the life system
type System struct {
	health  entity
	changed bool
	elapsed float32
}

// New is called when the system is added to the world
func (s *System) New(w *ecs.World) {
	empty := healthbar{BasicEntity: ecs.NewBasic()}
	emptys, _ := common.LoadedSprite("emptyHealth.png")
	empty.RenderComponent.Drawable = emptys
	empty.RenderComponent.SetZIndex(10)
	empty.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 10 + options.XOffset,
			Y: 10 + options.YOffset,
		},
		Width:  empty.RenderComponent.Drawable.Width(),
		Height: empty.RenderComponent.Drawable.Height(),
	}
	w.AddEntity(&empty)

	b := healthbar{BasicEntity: ecs.NewBasic()}
	bs, _ := common.LoadedSprite("health.png")
	b.RenderComponent.Drawable = bs
	b.RenderComponent.SetZIndex(11)
	b.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 10 + options.XOffset,
			Y: 10 + options.YOffset,
		},
		Width:  empty.RenderComponent.Drawable.Width(),
		Height: empty.RenderComponent.Drawable.Height(),
	}
	b.Component.Health = 100
	b.Component.Max = 100
	w.AddEntity(&b)

	s.health = entity{&b.BasicEntity, &b.RenderComponent, &b.SpaceComponent, &b.Component}

	hitp, _ := common.LoadedPlayer("hit.wav")
	hitp.SetVolume(options.TheOptions.SFXLevel)
	hit := struct {
		ecs.BasicEntity
		common.AudioComponent
	}{BasicEntity: ecs.NewBasic()}
	hit.AudioComponent.Player = hitp
	w.AddEntity(&hit)

	engo.Mailbox.Listen(messages.DamageType, func(m engo.Message) {
		msg, ok := m.(messages.Damage)
		if !ok {
			return
		}
		if msg.Amount > 0 {
			s.changed = true
			s.health.Component.Health += msg.Amount
			return
		}
		if s.elapsed <= 0 {
			s.changed = true
			s.health.Component.Health += msg.Amount
			if options.TheOptions.SFX {
				hitp.Rewind()
				hitp.Play()
			}
			engo.Mailbox.Dispatch(messages.Flash{})
		}
	})
}

// Remove doesn't do anything since the only entity is the healthbar
func (s *System) Remove(basic ecs.BasicEntity) {}

// Update is called once per frame
func (s *System) Update(dt float32) {
	if s.elapsed > 0.3 {
		s.elapsed = 0
		return
	}
	if s.elapsed > 0 {
		s.elapsed += dt
		return
	}
	if s.changed {
		s.elapsed += dt
		if s.health.Health <= 0 {
			s.changed = false
			engo.Mailbox.Dispatch(messages.GameOver{})
			return
		}
		if s.health.Health > 100 {
			engo.Mailbox.Dispatch(messages.Score{Amount: int(s.health.Health-100) * 5})
			s.health.Health = 100
		}
		s.health.RenderComponent.Scale.X = s.health.Health / s.health.Max
		s.changed = false
	}
}
