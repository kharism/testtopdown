package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image  *ebiten.Image
	Facing int
}

var Sprite = donburi.NewComponentType[SpriteData]()
