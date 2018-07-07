package bullet

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*Component
}

type bullet struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	common.AnimationComponent
	common.CollisionComponent
	Component
}

type sound struct {
	ecs.BasicEntity
	common.AudioComponent
}
