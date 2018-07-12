package musicdown

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.MouseComponent
	*common.AudioComponent
}
