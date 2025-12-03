package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kharism/hanashi/core"
	"github.com/kharism/testtopdown/assets"
)

var hanashiScene *core.Scene

func isSpaceBarPressed() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace)
}

type basicLayouter struct{}

// return width and height of the scene
func (*basicLayouter) GetLayout() (width, height int) {
	return 64 * 15, 64 * 10
}

// return the starting text position where the box containing name of the character appear on the scene
// return negative number if no such box needed
func (*basicLayouter) GetNamePosition() (x, y int) {
	return 10, 640 - 120
}

// get the starting position of the text
func (*basicLayouter) GetTextPosition() (x, y int) {
	return 10, 640 - 110
}
func GetHanashiScene() *core.Scene {
	if hanashiScene != nil {
		hanashiScene.Events = []core.Event{}
		hanashiScene.EventIndex = 0
		return hanashiScene
	}
	hanashiScene = core.NewScene()
	hanashiScene.Events = []core.Event{}
	core.DetectKeyboardNext = isSpaceBarPressed
	hanashiScene.FontFace = assets.Face
	hanashiScene.TxtBg = ebiten.NewImage(1024-128, 128)
	hanashiScene.TxtBg.Fill(color.RGBA{R: 0x4f, G: 0x8f, B: 0xba, A: 255})

	hanashiScene.SetLayouter(&basicLayouter{})
	return hanashiScene
}
