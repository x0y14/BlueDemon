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
	Player       *Object
	Objs         map[string]*Object
	Enemies      map[string]*Object
	currentStage *Stage
	Stages       map[string]*Stage
}

func (g *Game) SetStage(id string) {
	g.currentStage = g.Stages[id]
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
		float64(g.Player.Center.X+g.Player.SkinVertex.X),
		float64(g.Player.Center.Y+g.Player.SkinVertex.Y))
	screen.DrawImage(g.Player.Skin, op)
}

func (g *Game) drawObject(screen *ebiten.Image, obj *Object) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(obj.Center.X+obj.SkinVertex.X),
		float64(obj.Center.Y+obj.SkinVertex.Y))
	screen.DrawImage(obj.Skin, op)
}

func (g *Game) drawStage(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// ベースとなるスキンを貼る
	screen.DrawImage(g.currentStage.Skin, op)
	// 小物を貼る
	for _, obj := range g.currentStage.Objs {
		g.drawObject(screen, obj)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawStage(screen)
	g.drawPlayer(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
