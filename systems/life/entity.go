package life

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
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
