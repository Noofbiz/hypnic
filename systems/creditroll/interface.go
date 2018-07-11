package creditroll

import "engo.io/engo/common"

// Face is the interface for an entity with a bullet component
type Face interface {
	GetCreditRollComponent() *Component
}

// Able is the interface for if the entity is compatible with the bullet system
type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
