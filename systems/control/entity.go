package control

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*Component
}

type anim struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*common.RenderComponent
	*common.AnimationComponent
}
