package life

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.RenderComponent
	*common.SpaceComponent
	*Component
}

type healthbar struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	Component
}
