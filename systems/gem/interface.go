package gem

import "github.com/EngoEngine/engo/common"

type Face interface {
	GetGemComponent() *Component
}

type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
