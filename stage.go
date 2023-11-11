package blueDemon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

type Stage struct {
	Id   string
	Objs []*Object
	Skin *ebiten.Image
}

func NewStage(id string, objs []*Object, skin *ebiten.Image) *Stage {
	return &Stage{
		Id:   id,
		Objs: objs,
		Skin: skin,
	}
}

var Stage1 *Stage

func init() {
	stageSkin, _, err := ebitenutil.NewImageFromFile("assets/empty.png")
	if err != nil {
		log.Fatalln(err)
	}
	tileSkin, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileBody := NewRectanglePolygon(NewPosition(0, 0), 80, 80, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x80})
	tile := NewObject("tile", []*Polygon{tileBody}, NewPosition(0, 0), tileSkin, NewPosition(-tileBody.Width/2, -tileBody.Height/2))

	Stage1 = NewStage("stage1", []*Object{tile}, stageSkin)
}
