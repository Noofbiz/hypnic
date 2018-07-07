package score

import (
	"engo.io/ecs"
	"engo.io/engo/common"
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
