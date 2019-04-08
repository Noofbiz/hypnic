package bullet

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
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
