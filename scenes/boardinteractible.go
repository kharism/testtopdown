package scenes

import (
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/donburi/ecs"
)

type BoardObject struct {
	*tiled.Object

	newTiledName string

	targetCol int
	targetRow int

	scene *TiledScenes
}

func (d *BoardObject) Interact(ecs *ecs.ECS) {

}
