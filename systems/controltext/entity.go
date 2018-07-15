package controltext

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*common.RenderComponent
}
