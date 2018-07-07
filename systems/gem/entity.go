package gem

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*Component
}

type gem struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	common.CollisionComponent
	Component
}

type sound struct {
	ecs.BasicEntity
	common.AudioComponent
}
