package assets

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"os"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed basic_character_all.png
var player []byte

var PlayerImg *ebiten.Image

//go:embed PixelOperator8.ttf
var PixelFontTTF []byte
var PixelFont *text.GoTextFaceSource
var Face *text.GoTextFace

func init() {
	var err error
	PlayerImg, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(player))

	if err != nil {
		fmt.Printf("Error loading player spirte: %s", err.Error())
		os.Exit(2)
	}
	s, err := text.NewGoTextFaceSource(bytes.NewReader(PixelFontTTF))
	if err != nil {
		log.Fatal(err)
	}
	PixelFont = s
	Face = &text.GoTextFace{
		Source: PixelFont,
		Size:   15,
	}
}

// assuming 64*64 per frame and 2 frame animation
func LoadPlayerSubSprite(facing, frame int) *ebiten.Image {
	yStart := facing * 64
	xStart := frame * 64

	xEnd := xStart + 64
	yEnd := yStart + 64

	return PlayerImg.SubImage(image.Rect(xStart, yStart, xEnd, yEnd)).(*ebiten.Image)
}

const PLAYER_FACE_FRONT = 0
const PLAYER_FACE_LEFT = 1
const PLAYER_FACE_RIGHT = 2
const PLAYER_FACE_BACK = 3
