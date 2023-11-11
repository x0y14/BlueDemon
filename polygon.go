package blueDemon

import "image/color"

type Polygon struct {
	Figure
	Center *Position
	Width  float32
	Height float32
	Radius float32
	Color  color.Color
}
type Figure int

const (
	Rectangle Figure = iota
)

func NewRectanglePolygon(center *Position, w, h float32, c color.Color) *Polygon {
	return &Polygon{
		Figure: Rectangle,
		Center: center,
		Width:  w,
		Height: h,
		Radius: 0,
		Color:  c,
	}
}
