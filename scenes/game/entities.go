package game

import (
	"engo.io/ecs"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/systems/control"
	"github.com/Noofbiz/hypnic/systems/wall"
)

type background struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type wallTile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	wall.Component
}

type player struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	control.Component
}
