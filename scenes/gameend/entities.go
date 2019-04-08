package gameend

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type sprite struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type button struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}
