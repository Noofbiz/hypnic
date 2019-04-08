package potion

import "github.com/EngoEngine/engo/common"

type Face interface {
	GetPotionComponent() *Component
}

type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
