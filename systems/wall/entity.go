package wall

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type entity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*Component
}

type entityList []entity

func (e entityList) Len() int {
	return len(e)
}

func (e entityList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e entityList) Less(i, j int) bool {
	return e[i].SpaceComponent.Position.Y < e[j].SpaceComponent.Position.Y
}
