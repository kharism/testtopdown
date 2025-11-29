package scenes

import (
	"image"

	"github.com/joelschutz/stagehand"
)

type BaseScene struct {
	bounds image.Rectangle
	sm     *stagehand.SceneManager[SceneData]
}

func (s *BaseScene) Layout(w, h int) (int, int) {
	s.bounds = image.Rect(0, 0, w, h)
	return w, h
}
