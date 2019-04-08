package bullet

import "github.com/EngoEngine/engo/common"

// Face is the interface for an entity with a bullet component
type Face interface {
	GetBulletComponent() *Component
}

// Able is the interface for if the entity is compatible with the bullet system
type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
