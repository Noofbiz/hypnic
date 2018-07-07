package potion

import "engo.io/engo/common"

type Face interface {
	GetPotionComponent() *Component
}

type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
