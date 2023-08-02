package gamecore

import (
	"log"
	
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	control *control
}

func createGame(control *control) *game {
	return &game{
		control: control,
	}
}

func (g *game) Update() error {
	defer g.catch()
	for _, m := range g.control.gamecore.modules {
		m.updateComponents()
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.control.screen = screen
	for _, m := range g.control.gamecore.modules {
		m.renderComponents()
	}
}

func (g *game) Layout(ww, wh int) (int, int) {
	g.control.gamecore.control.windowManager.width = ww
	g.control.gamecore.control.windowManager.height = wh
	return ww, wh
}

func (g *game) catch() {
	if r := recover(); r != nil {
		log.Fatalln(r)
	}
}
