package flash

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/Noofbiz/hypnic/messages"
)

// System makes the player flash when damaged
type System struct {
	e     entity
	blink bool
}

// New is called when the system is added to the world
func (s *System) New(w *ecs.World) {
	engo.Mailbox.Listen(messages.FlashType, func(m engo.Message) {
		_, ok := m.(messages.Flash)
		if !ok {
			return
		}
		s.blink = true
	})
}

// Add adds an entity to the system
func (s *System) Add(basic *ecs.BasicEntity, anim *common.AnimationComponent, flash *FComponent) {
	s.e = entity{basic, anim, flash}
}

// AddByInterface adds an entity to the system that implements the Able interface
func (s *System) AddByInterface(i ecs.Identifier) {
	o, _ := i.(Able)
	s.Add(o.GetBasicEntity(), o.GetAnimationComponent(), o.GetFlashComponent())
}

// Remove removes an entity from the system
func (s *System) Remove(basic ecs.BasicEntity) {}

// Update is called once per frame
func (s *System) Update(dt float32) {
	if s.blink {
		s.e.AnimationComponent.SelectAnimationByName("flash")
		s.blink = false
	}
}
