package credits

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"

	"github.com/Noofbiz/hypnic/systems/creditroll"
)

type sprite struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type mobilesprite struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	creditroll.Component
}

type bgm struct {
	ecs.BasicEntity
	common.AudioComponent
}
