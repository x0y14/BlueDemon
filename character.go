package blueDemon

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

var Human *Object

func init() {
	body := NewRectanglePolygon(
		NewPosition(0, 0),
		60,
		120,
		color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x80},
	)

	skin, _, err := ebitenutil.NewImageFromFile("assets/human.png")
	if err != nil {
		log.Fatalln(err)
	}
	Human = NewObject(
		"Player",
		[]*Polygon{body},
		nil,
		skin,
		NewPosition(-body.Width/2, -body.Height/2),
	)
}
