package scenes

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/joelschutz/stagehand"
	"github.com/kharism/hanashi/core"
	"github.com/kharism/testtopdown/assets"
	"github.com/kharism/testtopdown/scenes/components"
	"github.com/kharism/testtopdown/scenes/system"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	//"github.com/lafriks/go-tiled/render"
)

type TiledScenes struct {
	BaseScene
	bg          *ebiten.Image
	tiledLayout *tiled.Map
	//sm          *stagehand.SceneManager[SceneData]
	sceneData SceneData
	ecs       *ecs.ECS
	player    *donburi.Entry
	Scene     *core.Scene

	tileRenderer   *system.BgRenderer
	spriteRenderer *system.SpriteRenderer
	movementSystem *system.PlayerMovement
}

func (tiledScenes *TiledScenes) Update() error {
	if tiledScenes.Scene == nil {
		tiledScenes.ecs.Update()
	} else {
		tiledScenes.Scene.Update()
	}
	return nil
}

func (tiledScenes *TiledScenes) Draw(screen *ebiten.Image) {
	// base := ebiten.NewImage(TileWidth*WidthInTile, TileHeight*HeightInTile)
	tiledScenes.ecs.DrawLayer(0, screen)
	tiledScenes.ecs.DrawLayer(1, screen)
	logPos := components.LogicalPos.Get(tiledScenes.player)
	dbgTxt := fmt.Sprintf("PlayerPos %f %f\nPlayer Speed %f %f\nFPS: %f", logPos.X, logPos.Y, system.SpeedX, system.SpeedY, ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, dbgTxt)
	if tiledScenes.Scene != nil {
		tiledScenes.Scene.Draw(screen)
	}
	// screen.DrawImage(screen, &ebiten.DrawImageOptions{})
	return
}
func createPlayerEntry(ecs *ecs.ECS) *donburi.Entry {
	player := ecs.World.Create(components.GridPos, components.Sprite, components.LogicalPos)
	playerEntry := ecs.World.Entry(player)
	gridPos := components.GridPos.Get(playerEntry)
	gridPos.Col = 4
	gridPos.Row = 4

	//components.Sprite.Set(playerEntry, PlayerImg)
	return playerEntry
}

// get static interactibles
func (s *TiledScenes) generateInteractiblesMap(tiledMap *tiled.Map) []system.Interactible {
	objectsTile := make([]system.Interactible, len(tiledMap.Layers[0].Tiles))
	for _, o := range tiledMap.ObjectGroups[0].Objects {
		col := int(o.X) / TileWidth
		row := int(o.Y) / TileHeight
		realIdx := row*tiledMap.Width + col
		if o.Type == "exitpoint" {
			k := &DoorObject{
				Object:       o,
				newTiledName: o.Properties.GetString("NextScene"),
				targetCol:    o.Properties.GetInt("NextSceneLocationX"),
				targetRow:    o.Properties.GetInt("NextSceneLocationY"),
				scene:        s,
			}
			k.targetCol = o.Properties.GetInt("NextSceneLocationX")
			objectsTile[realIdx] = k
		} else if o.Type == "board" {

		}

	}
	return objectsTile
}

func (s *TiledScenes) changeTileLevel(levelName string) {
	gameMap, err := tiled.LoadFile(levelName)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}
	err = renderer.RenderLayer(0)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Get a reference to the Renderer's output, an image.NRGBA struct.
	img := renderer.Result
	s.bg = ebiten.NewImageFromImage(img)
	err = renderer.RenderLayer(1)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}
	img2 := ebiten.NewImageFromImage(renderer.Result)
	s.tiledLayout = gameMap
	s.tileRenderer.Bg1 = s.bg
	s.tileRenderer.Bg2 = img2
	objMap := s.generateInteractiblesMap(gameMap)

	s.movementSystem.LegalMoveList = assets.AllowedTiles[gameMap.Tilesets[0].Name]
	s.movementSystem.Map = gameMap.Layers[0].Tiles
	s.movementSystem.MapWidth = gameMap.Width
	s.movementSystem.ObjectMap = objMap

}

var PLAYER_FACING = 0

func (s *TiledScenes) Load(state SceneData, manager stagehand.SceneController[SceneData]) {
	// your load code

	s.sm = manager.(*stagehand.SceneManager[SceneData]) // This type assertion is important
	world := donburi.NewWorld()
	s.ecs = ecs.NewECS(world)
	s.sceneData = state
	gameMap, err := tiled.LoadFile(state.MapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	s.tiledLayout = gameMap
	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}
	playerEntry := createPlayerEntry(s.ecs)
	playerGP := components.GridPos.Get(playerEntry)
	playerGP.Col = s.sceneData.PlayerGridPosX
	playerGP.Row = s.sceneData.PlayerGridPosY

	playerLP := components.LogicalPos.Get(playerEntry)
	playerLP.X = TileWidth * float64(playerGP.Col)
	playerLP.Y = TileHeight * float64(playerGP.Row)

	s.player = playerEntry
	if components.Sprite.Get(playerEntry).Image == nil {
		components.Sprite.Set(playerEntry, &components.SpriteData{
			Image:  assets.LoadPlayerSubSprite(0, 0),
			Facing: 0,
		})
	}

	err = renderer.RenderLayer(0)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Get a reference to the Renderer's output, an image.NRGBA struct.
	img := renderer.Result

	s.bg = ebiten.NewImageFromImage(img)
	err = renderer.RenderLayer(1)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}
	img2 := ebiten.NewImageFromImage(renderer.Result)

	tileRenderer := &system.BgRenderer{
		Player:             playerEntry,
		Bg1:                s.bg,
		Bg2:                img2,
		ScreenWidthInTile:  15,
		ScreenHeightInTile: 10,
	}
	spriteRenderer := &system.SpriteRenderer{
		Player: playerEntry,
		Query: donburi.NewQuery(
			filter.Contains(
				components.Sprite,
				components.GridPos,
			),
		),
		GameWidthInTile:  WidthInTile,
		GameHeightInTile: HeightInTile,
		TileHeight:       TileHeight,
		TileWidth:        TileWidth,
	}
	objMap := s.generateInteractiblesMap(gameMap)
	playerMovement := &system.PlayerMovement{
		Player:        playerEntry,
		LegalMoveList: assets.AllowedTiles[gameMap.Tilesets[0].Name],
		Map:           gameMap.Layers[0].Tiles,
		MapWidth:      gameMap.Width,
		ObjectMap:     objMap,
	}
	s.movementSystem = playerMovement
	s.spriteRenderer = spriteRenderer
	s.tileRenderer = tileRenderer
	s.ecs.AddRenderer(0, tileRenderer.RenderBg)
	s.ecs.AddRenderer(1, spriteRenderer.RenderSprite)
	s.ecs.AddSystem(playerMovement.Update)

}

func (s *TiledScenes) Unload() SceneData {
	// your unload code
	return s.sceneData
}
