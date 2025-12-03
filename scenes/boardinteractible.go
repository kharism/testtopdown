package scenes

import (
	"github.com/kharism/hanashi/core"
	"github.com/kharism/testtopdown/assets"
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/donburi/ecs"
)

type BoardObject struct {
	*tiled.Object

	Texts []string

	targetCol int
	targetRow int

	scene *TiledScenes
}

func (d *BoardObject) Interact(ecs *ecs.ECS) {
	hanashiScene := GetHanashiScene()
	events := []core.Event{}
	for _, l := range d.Texts {
		j := core.NewDialogueEvent("", l, assets.Face)
		events = append(events, j)
	}
	d.scene.Scene = hanashiScene
	hanashiScene.Events = events

	hanashiScene.Events[0].Execute(hanashiScene)
}
