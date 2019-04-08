package control

import "github.com/EngoEngine/engo/common"

// Face is an interface to get the control component
type Face interface {
	GetControlComponent() *Component
}

// Able is if it is able to be used by the control system
type Able interface {
	common.BasicFace
	common.SpaceFace
	Face
}
