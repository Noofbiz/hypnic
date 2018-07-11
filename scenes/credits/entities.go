package credits

import (
	"engo.io/ecs"
	"engo.io/engo/common"

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
