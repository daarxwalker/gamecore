package gamecore

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text"
)

type TextRenderer interface {
	Color(color color.Color) TextRenderer
	X(x int) TextRenderer
	Y(y int) TextRenderer
	Value(value any) TextRenderer
	Render()
}

type textRenderer struct {
	control *control
	color   color.Color
	value   string
	x       int
	y       int
}

func createTextRenderer(control *control) *textRenderer {
	return &textRenderer{
		control: control,
	}
}

func (r *textRenderer) Color(color color.Color) TextRenderer {
	r.color = color
	return r
}

func (r *textRenderer) X(x int) TextRenderer {
	r.x = x
	return r
}

func (r *textRenderer) Y(y int) TextRenderer {
	r.y = y + r.control.fontManager.getDefault().Metrics().Height.Ceil()
	return r
}

func (r *textRenderer) Render() {
	if r.color == nil {
		r.color = color.White
	}
	text.Draw(r.control.screen, r.value, r.control.fontManager.getDefault(), r.x, r.y, r.color)
}

func (r *textRenderer) Value(value any) TextRenderer {
	r.value = fmt.Sprintf("%v", value)
	return r
}
