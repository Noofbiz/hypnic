package control

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
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
