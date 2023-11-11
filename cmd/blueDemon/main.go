package main

import (
	"blueDemon"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(blueDemon.ScreenWidth, blueDemon.ScreenHeight)
	ebiten.SetWindowTitle("BlueDemon")

	game := &blueDemon.Game{}
	game.Player = blueDemon.Human
	game.Player.SetPosition(blueDemon.NewPosition(blueDemon.ScreenWidth/2, blueDemon.ScreenHeight/2))
	game.Stages = map[string]*blueDemon.Stage{
		"stage1": blueDemon.Stage1,
	}
	game.SetStage("stage1")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
