package gameend

import (
	"engo.io/ecs"
	"engo.io/engo/common"
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
	common.MouseComponent
}
