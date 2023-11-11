package blueDemon

import (
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 640
)

type Game struct {
	Player  *Object
	Objs    []*Object
	Enemies []*Object
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.SetPosition(
			NewPosition(g.Player.Center.X, g.Player.Center.Y-5))
	} // up
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.SetPosition(
			NewPosition(g.Player.Center.X, g.Player.Center.Y+5))
	} // down
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.SetPosition(
			NewPosition(g.Player.Center.X+5, g.Player.Center.Y))
	} // right
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.SetPosition(
			NewPosition(g.Player.Center.X-5, g.Player.Center.Y))
	} // left
	return nil
}

func (g *Game) drawPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(g.Player.Center.X-g.Player.Component[0].Width/2),
		float64(g.Player.Center.Y-g.Player.Component[0].Height/2))
	screen.DrawImage(g.Player.Skin, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawPlayer(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
