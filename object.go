package blueDemon

import "github.com/hajimehoshi/ebiten/v2"

type Object struct {
	Id        string
	Component []*Polygon
	Center    *Position

	currentState string
	States       map[string]*State
}
type State struct {
	Skin       *ebiten.Image
	SkinVertex *Position
}

func NewState(skin *ebiten.Image, vertex *Position) *State {
	return &State{
		Skin:       skin,
		SkinVertex: vertex,
	}
}

func NewObject(id string, component []*Polygon, center *Position, defaultState *State) *Object {
	return &Object{
		Id:           id,
		Component:    component,
		Center:       center,
		currentState: "default",
		States:       map[string]*State{"default": defaultState},
	}
}

func (o *Object) SetPosition(p *Position) {
	o.Center = p
}

func (o *Object) IsCollision(objs []*Object) (bool, string) { // (yes,no), id
	return false, ""
}

func (o *Object) CurrentState() string {
	return o.currentState
}

func (o *Object) UpdateState(newState string) {
	o.currentState = newState
}
