package main

import (
	"blueDemon"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(blueDemon.ScreenWidth, blueDemon.ScreenHeight)
	ebiten.SetWindowTitle("BlueDemon")
	if err := ebiten.RunGame(&blueDemon.Game{}); err != nil {
		panic(err)
	}
}
