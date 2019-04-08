package gem

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
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
