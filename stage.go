package blueDemon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"strconv"
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

	// 壁
	wallSkin, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
	if err != nil {
		log.Fatalln(err)
	}
	doorSkin, _, err := ebitenutil.NewImageFromFile("assets/door.png")
	if err != nil {
		log.Fatalln(err)
	}
	body80x160 := *NewRectanglePolygon(NewPosition(0, 0), 80, 160, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x80})
	wallState := *NewState(wallSkin, NewPosition(-body80x160.Width/2, -body80x160.Height/2))
	var walls []*Object
	for i := 0; i < 8; i++ {
		walls = append(walls,
			NewObject("wall-"+strconv.Itoa(i), []*Polygon{&body80x160}, NewPosition(40+float32(80*i), +80), &wallState),
		)
	}
	door := NewObject("door", []*Polygon{&body80x160}, NewPosition(walls[4].Center.X, walls[4].Center.Y), NewState(doorSkin, NewPosition(-body80x160.Width/2, -body80x160.Height/2)))
	walls[4] = door

	floorSkin, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
	body80x80 := *NewRectanglePolygon(NewPosition(0, 0), 80, 80, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x80})
	floorState := *NewState(floorSkin, NewPosition(-body80x80.Width/2, -body80x80.Height/2))
	var floors []*Object
	for i := 0; i < 8; i++ { //横
		for j := 0; j < 6; j++ { //たて
			floors = append(floors,
				NewObject("floor-"+strconv.Itoa(i)+"-"+strconv.Itoa(j), []*Polygon{&body80x80}, NewPosition(40+float32(80*i), 160+40+float32(80*j)), &floorState))
		}
	}

	var all []*Object
	all = append(all, walls...)
	all = append(all, floors...)
	Stage1 = NewStage("stage1", all, stageSkin)
}
