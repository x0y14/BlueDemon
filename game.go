package blueDemon

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 640
	ScreenHeight = 640
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
