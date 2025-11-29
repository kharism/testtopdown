package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kharism/testtopdown/assets"
	"github.com/kharism/testtopdown/scenes/components"
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type PlayerMovement struct {
	Player *donburi.Entry

	Map []*tiled.LayerTile

	// static interactible stored here
	ObjectMap []Interactible
	// width and height in tile
	MapHeight int
	MapWidth  int

	// grid value for legal move
	LegalMoveList []int

	keepMoving bool
}

type PlayerMoveListener interface {
	OnPlayerMove()
}

func (p *PlayerMovement) checkLegal(col, row int) (bool, Interactible) {
	realIdx := row*p.MapWidth + col
	if p.ObjectMap[realIdx] != nil {
		return false, p.ObjectMap[realIdx]
	}
	for _, v := range p.LegalMoveList {
		if p.Map[realIdx].ID == uint32(v) {
			return true, nil
		}
	}
	return false, nil
}
func (p *PlayerMovement) GetInteractible(col, row int) (bool, Interactible) {
	realIdx := row*p.MapWidth + col
	if p.ObjectMap[realIdx] != nil {
		return false, nil
	} else {
		return true, p.ObjectMap[realIdx]
	}
}

// movement stuff
var SpeedX = 0.0
var SpeedY = 0.0

//var CurrentFacing = 0

func (p *PlayerMovement) Update(ecs *ecs.ECS) {
	logicalPos := components.LogicalPos.Get(p.Player)
	spriteData := components.Sprite.Get(p.Player)
	gridPos := components.GridPos.Get(p.Player)
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		p.keepMoving = false
		if int(logicalPos.X)%64 == 0 {
			//spriteData.Image = assets.LoadPlayerSubSprite(spriteData.Facing, 1)
			SpeedX = 0
		}
		if int(logicalPos.Y)%64 == 0 {
			//spriteData.Image = assets.LoadPlayerSubSprite(spriteData.Facing, 1)
			SpeedY = 0
		}
		// SpeedX = 0
		// SpeedY = 0
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		ok, interactible := p.checkLegal(gridPos.Col, gridPos.Row+1)

		if !ok {
			spriteData.Image = assets.LoadPlayerSubSprite(assets.PLAYER_FACE_FRONT, 1)
			spriteData.Facing = assets.PLAYER_FACE_FRONT
			if interactible != nil {
				interactible.Interact(ecs)
				return
			}
			p.keepMoving = false
			SpeedY = 0.0

			return
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			spriteData.Image = assets.LoadPlayerSubSprite(assets.PLAYER_FACE_FRONT, 1)
			spriteData.Facing = assets.PLAYER_FACE_FRONT
			SpeedY = 8.0
			SpeedX = 0
		} else {
			p.keepMoving = true

		}
		//CurrentFacing = assets.PLAYER_FACE_FRONT

	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		ok, interactible := p.checkLegal(gridPos.Col-1, gridPos.Row)

		if !ok {
			spriteData.Image = assets.LoadPlayerSubSprite(assets.PLAYER_FACE_LEFT, 1)
			spriteData.Facing = assets.PLAYER_FACE_LEFT
			if interactible != nil {
				interactible.Interact(ecs)
				return
			}
			p.keepMoving = false
			SpeedY = 0.0

			return
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			spriteData.Facing = assets.PLAYER_FACE_LEFT
			SpeedX = -8.0
			SpeedY = 0
		} else {
			p.keepMoving = true
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		ok, interactible := p.checkLegal(gridPos.Col, gridPos.Row-1)

		if !ok {
			spriteData.Image = assets.LoadPlayerSubSprite(assets.PLAYER_FACE_BACK, 1)
			spriteData.Facing = assets.PLAYER_FACE_BACK

			if interactible != nil {
				interactible.Interact(ecs)
				return
			}
			p.keepMoving = false
			SpeedY = 0.0
			return
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			spriteData.Facing = assets.PLAYER_FACE_BACK
			SpeedY = -8.0
			SpeedX = 0
		} else {
			p.keepMoving = true
			// if int(logicalPos.X)%64 == 0 {
			// 	spriteData.Image = assets.LoadPlayerSubSprite(assets.PLAYER_FACE_BACK, 1)
			// }
		}
		//CurrentFacing = assets.PLAYER_FACE_BACK

	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		ok, interactible := p.checkLegal(gridPos.Col+1, gridPos.Row)

		if !ok {
			spriteData.Image = assets.LoadPlayerSubSprite(assets.PLAYER_FACE_RIGHT, 1)
			spriteData.Facing = assets.PLAYER_FACE_RIGHT

			if interactible != nil {
				interactible.Interact(ecs)
				return
			}
			p.keepMoving = false
			SpeedY = 0.0
			return
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			spriteData.Facing = assets.PLAYER_FACE_RIGHT
			SpeedX = 8.0
			SpeedY = 0
		} else {
			p.keepMoving = true
		}

	}
	if p.keepMoving {
		if int(logicalPos.X)%64 == 0 {
			spriteData.Image = assets.LoadPlayerSubSprite(spriteData.Facing, 1)
		}
		if int(logicalPos.Y)%64 == 0 {
			spriteData.Image = assets.LoadPlayerSubSprite(spriteData.Facing, 1)
		}
	}
	if SpeedX != 0 {
		logicalPos.X += SpeedX
	} else {
		logicalPos.Y += SpeedY
	}

	// if int(logicalPos.X)%64 == 32 || int(logicalPos.Y)%64 == 32 {
	// 	components.Sprite.Set(p.Player, assets.LoadPlayerSubSprite(CurrentFacing, 0))
	// }
	if SpeedX != 0 && int(logicalPos.X)%64 == 0 {
		gridPos.Col = int(logicalPos.X / 64)
		if !p.keepMoving {
			SpeedX = 0

			spriteData := components.Sprite.Get(p.Player)
			spriteData.Image = assets.LoadPlayerSubSprite(spriteData.Facing, 0)
		} else {
			spriteData := components.Sprite.Get(p.Player)
			spriteData.Image = assets.LoadPlayerSubSprite(spriteData.Facing, 0)
		}
	}
	if SpeedY != 0 && int(logicalPos.Y)%64 == 0 {
		gridPos.Row = int(logicalPos.Y / 64)
		if !p.keepMoving {
			SpeedY = 0

			spriteData := components.Sprite.Get(p.Player)
			spriteData.Image = assets.LoadPlayerSubSprite(spriteData.Facing, 0)
		} else {
			spriteData := components.Sprite.Get(p.Player)
			spriteData.Image = assets.LoadPlayerSubSprite(spriteData.Facing, 0)
		}

	}
	return
}
