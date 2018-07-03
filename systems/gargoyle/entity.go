package gargoyle

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*Component
}

type gargoyle struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	Component
}
