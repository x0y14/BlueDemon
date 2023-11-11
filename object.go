package blueDemon

import "github.com/hajimehoshi/ebiten/v2"

type Object struct {
	Id        string
	Component []*Polygon
	Center    *Position
	Skin      *ebiten.Image
}

func NewObject(id string, component []*Polygon, center *Position, skin *ebiten.Image) *Object {
	return &Object{
		Id:        id,
		Component: component,
		Center:    center,
		Skin:      skin,
	}
}

func (o *Object) SetPosition(p *Position) {
	o.Center = p
}

func (o *Object) IsCollision(objs []*Object) (bool, string) { // (yes,no), id
	return false, ""
}
