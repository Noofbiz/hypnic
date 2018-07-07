package flash

import "engo.io/engo/common"

// Face is the flash component's interface
type Face interface {
	GetFlashComponent() *FComponent
}

// Able is an entity that's able to be in the flash system
type Able interface {
	common.BasicFace
	common.AnimationFace
	Face
}
