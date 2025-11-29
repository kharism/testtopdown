package scenes

import (
	"github.com/kharism/testtopdown/scenes/components"
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/donburi/ecs"
)

type DoorObject struct {
	*tiled.Object

	newTiledName string

	targetCol int
	targetRow int

	scene *TiledScenes
}

func (d *DoorObject) Interact(ecs *ecs.ECS) {
	gp := components.GridPos.Get(d.scene.player)
	gp.Col = d.targetCol
	gp.Row = d.targetRow

	lp := components.LogicalPos.Get(d.scene.player)
	lp.X = float64(TileWidth * d.targetCol)
	lp.Y = float64(TileHeight * d.targetRow)

	d.scene.changeTileLevel(d.newTiledName)
}
