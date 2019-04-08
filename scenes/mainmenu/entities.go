package mainmenu

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
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
	common.AudioComponent
}

type bgm struct {
	ecs.BasicEntity
	common.AudioComponent
}
