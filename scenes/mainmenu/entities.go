package mainmenu

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type background struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type label struct {
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

type bgm struct {
	ecs.BasicEntity
	common.AudioComponent
}
