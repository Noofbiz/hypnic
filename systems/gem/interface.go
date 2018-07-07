package gem

import "engo.io/engo/common"

type Face interface {
	GetGemComponent() *Component
}

type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
