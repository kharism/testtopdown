package system

import (
	"github.com/yohamta/donburi/ecs"
)

type Interactible interface {
	Interact(ecs *ecs.ECS)
}
