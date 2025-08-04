package scenes

import "github.com/hajimehoshi/ebiten/v2"

type SceneId uint

const (
	GameScenes SceneId = iota
	StartSceneId
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
	FirstLoad()
	OnEnter()
	OnExit()
}
