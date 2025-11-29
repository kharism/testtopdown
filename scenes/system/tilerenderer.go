package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/testtopdown/scenes/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type BgRenderer struct {
	Player *donburi.Entry
	//MapRenderer        *render.Renderer
	Bg1                *ebiten.Image
	Bg2                *ebiten.Image
	ScreenWidthInTile  int
	ScreenHeightInTile int
}

// render background so that the player always in center
func (b *BgRenderer) RenderBg(ecs *ecs.ECS, screen *ebiten.Image) {
	playerPos := components.LogicalPos.Get(b.Player)
	//b.MapRenderer.RenderLayer(0)
	ebiBg := b.Bg1
	//b.MapRenderer.RenderLayer(1)
	ebiBg2 := b.Bg2
	renderStartCol := playerPos.X - float64(b.ScreenWidthInTile/2*64)
	renderStartRow := playerPos.Y - float64(b.ScreenHeightInTile/2*64)
	geom := ebiten.GeoM{}
	geom.Translate(-renderStartCol, -renderStartRow)

	screen.DrawImage(ebiBg, &ebiten.DrawImageOptions{GeoM: geom})
	screen.DrawImage(ebiBg2, &ebiten.DrawImageOptions{GeoM: geom})
}
