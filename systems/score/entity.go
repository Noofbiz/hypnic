package score

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.RenderComponent
}

type text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}
