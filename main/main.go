package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/testtopdown/scenes"
)

func main() {
	scene1 := &scenes.TiledScenes{}
	state := &scenes.SceneData{
		MapPath:        "./assets/level1.tmx",
		PlayerGridPosX: 15,
		PlayerGridPosY: 8,
	}
	manager := stagehand.NewSceneManager(scene1, *state)

	ebiten.SetWindowSize(64*15, 64*10)

	if err := ebiten.RunGame(manager); err != nil {
		log.Fatal(err)
	}
}
