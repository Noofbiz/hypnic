package wall

import "github.com/EngoEngine/engo/common"

// Face is something able to retrieve it's wall Component
type Face interface {
	GetWallComponent() *Component
}

// Able is the interface for System
type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
