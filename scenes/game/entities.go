package game

import (
	"engo.io/ecs"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/systems/control"
	"github.com/Noofbiz/hypnic/systems/flash"
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
	common.AnimationComponent
	common.SpaceComponent
	common.RenderComponent
	common.CollisionComponent
	control.Component
	flash.FComponent
}

type shieldEntity struct {
	ecs.BasicEntity
	common.AnimationComponent
	common.RenderComponent
}

type bgm struct {
	ecs.BasicEntity
	common.AudioComponent
}
