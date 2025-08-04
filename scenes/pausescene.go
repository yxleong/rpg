package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct {
	loaded bool
}

func NewPauseScene() *PauseScene {
	return &PauseScene{
		loaded: false,
	}
}

func (s *PauseScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return GameSceneId
	}
	return PauseSceneId
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{117, 152, 146, 0})
	ebitenutil.DebugPrint(screen, "Press enter to unpause.\nPress Q to exit.")
}

func (s *PauseScene) FirstLoad() {
	s.loaded = true
}

func (s *PauseScene) OnEnter() {
}

func (s *PauseScene) OnExit() {
}

func (s *PauseScene) IsLoaded() bool {
	return s.loaded
}

var _ Scene = (*PauseScene)(nil)
