package control

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/controllers"
	"github.com/Noofbiz/hypnic/messages"
	"github.com/Noofbiz/hypnic/options"
)

// System is the system for controlling the player
type System struct {
	player entity
	speed  float32
	audio  *common.Player
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
			Position: s.player.SpaceComponent.Center(),
		})
	})
}

// Add adds the player to the system
func (s *System) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, control *Component) {
	s.player = entity{basic, space, control}
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
	s.player.Position.Add(engo.Point{
		X: s.speed * controlValue(),
		Y: 0,
	})
	if s.player.Position.X < 31+options.XOffset {
		engo.Mailbox.Dispatch(messages.Damage{
			Amount: -5,
		})
		s.player.Position.X = 32 + options.XOffset
	}
	if s.player.Position.X > 256+options.XOffset {
		engo.Mailbox.Dispatch(messages.Damage{
			Amount: -5,
		})
		s.player.Position.X = 255 + options.XOffset
	}
}

func controlValue() float32 {
	switch options.TheOptions.Controls {
	case "Acceler":
		return controllers.GetAccelerometerValue()
	case "Touch":
		return engo.Input.Axis(engo.DefaultMouseXAxis).Value()
	default:
		return engo.Input.Axis("movement").Value()
	}
}
