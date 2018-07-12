package musicbx

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.MouseComponent
	*common.RenderComponent
	*common.AudioComponent
}
