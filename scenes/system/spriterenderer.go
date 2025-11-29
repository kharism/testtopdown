package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/testtopdown/scenes/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type SpriteRenderer struct {
	Player *donburi.Entry
	Query  *donburi.Query

	TileWidth  int
	TileHeight int

	GameWidthInTile  int
	GameHeightInTile int
}

func (b *SpriteRenderer) RenderSprite(ecs *ecs.ECS, screen *ebiten.Image) {
	playerPos := components.GridPos.Get(b.Player)
	renderStartCol := float64(playerPos.Col - b.GameWidthInTile/2)
	renderStartRow := float64(playerPos.Row - b.GameHeightInTile/2)
	basicTranslateX := -float64(b.TileWidth) * renderStartCol
	basicTranslateY := -float64(b.TileHeight) * renderStartRow
	b.Query.Each(ecs.World, func(e *donburi.Entry) {
		gridPos := components.GridPos.Get(e)
		scrPosX := gridPos.Col * b.TileWidth
		scrPosY := gridPos.Row * b.TileHeight
		geom := ebiten.GeoM{}
		geom.Translate(float64(scrPosX)+basicTranslateX, float64(scrPosY)+basicTranslateY)
		s := components.Sprite.Get(b.Player)
		screen.DrawImage(s.Image, &ebiten.DrawImageOptions{GeoM: geom})
	})
}
