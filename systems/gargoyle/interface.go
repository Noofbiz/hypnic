package gargoyle

import "engo.io/engo/common"

// Face is an entity that is a gargoyle
type Face interface {
	GetGargoyleComponent() *Component
}

// Able is able to be added to the gargoyle system
type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
