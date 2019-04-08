package control

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/Noofbiz/hypnic/controllers"
	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

// System is the system for controlling the player
type System struct {
	entities []entity
	speed    float32
	audio    *common.Player
}

// New is called when the system is added to the world
func (s *System) New(w *ecs.World) {
	engo.Input.RegisterAxis(
		"movement",
		engo.AxisKeyPair{Min: engo.KeyArrowLeft, Max: engo.KeyArrowRight},
		engo.AxisKeyPair{Min: engo.KeyA, Max: engo.KeyD},
	)
	engo.Input.RegisterAxis(engo.DefaultMouseXAxis, engo.NewAxisMouse(engo.AxisMouseHori))
	s.speed = 3

	engo.Mailbox.Listen(messages.SendPlayerPositionType, func(m engo.Message) {
		_, ok := m.(messages.SendPlayerPosition)
		if !ok {
			return
		}
		engo.Mailbox.Dispatch(messages.GetPlayerPosition{
			Position: s.entities[0].SpaceComponent.Center(),
		})
	})
}

// Add adds the player to the system
func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, control *Component) {
	s.entities = append(s.entities, entity{basic, space, control})
}

// AddByInterface allows you to add a player as long as it implements the control interface
func (s *System) AddByInterface(i ecs.Identifier) {
	o, _ := i.(Able)
	s.Add(o.GetBasicEntity(), o.GetSpaceComponent(), o.GetControlComponent())
}

// Remove does nothing since there's only one player in the system
func (s *System) Remove(basic ecs.BasicEntity) {}

// Update is called once per frame
func (s *System) Update(dt float32) {
	dx := s.speed * controlValue()
	for _, e := range s.entities {
		e.Position.Add(engo.Point{
			X: dx,
			Y: 0,
		})
		if e.Position.X < 31+options.XOffset+e.XOff {
			engo.Mailbox.Dispatch(messages.Damage{
				Amount: -2.5,
			})
			e.Position.X = 32 + options.XOffset + e.XOff
		}
		if e.Position.X > 256+options.XOffset+e.XOff {
			engo.Mailbox.Dispatch(messages.Damage{
				Amount: -2.5,
			})
			e.Position.X = 255 + options.XOffset + e.XOff
		}
	}
}

func controlValue() float32 {
	switch options.TheOptions.Controls {
	case "Acceler":
		return controllers.GetAccelerometerValue()
	case "Touch":
		return engo.Input.Axis(engo.DefaultMouseXAxis).Value()
	default:
		return engo.Input.Axis("movement").Value() * 2
	}
}
