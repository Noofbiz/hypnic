package gargoyle

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
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
