package gamecore

import (
	"log"
	
	"github.com/hajimehoshi/ebiten/v2"
)

type Gamecore interface {
	Title(title string) Gamecore
	Window(w, h int) Gamecore
	Module(moduleFn func(Module)) Gamecore
	Run()
}

type gamecore struct {
	game    *game
	control *control
	modules []*module
	title   string
}

func New() Gamecore {
	gc := new(gamecore)
	gc.control = createControl(gc)
	gc.game = createGame(gc.control)
	return gc
}

func (c *gamecore) Module(moduleFn func(Module)) Gamecore {
	m := createModule(c.control)
	moduleFn(m)
	m.initComponents()
	c.modules = append(c.modules, m.getPtr())
	return c
}

func (c *gamecore) Title(title string) Gamecore {
	c.title = title
	return c
}

func (c *gamecore) Window(w, h int) Gamecore {
	c.control.windowManager = createWindowManager(w, h)
	return c
}

func (c *gamecore) Run() {
	c.initWindow()
	if err := ebiten.RunGame(
		c.game,
	); err != nil {
		log.Fatal(err)
	}
}

func (c *gamecore) initWindow() {
	if c.control.windowManager == nil {
		return
	}
	ebiten.SetWindowSize(c.control.windowManager.width, c.control.windowManager.height)
	ebiten.SetWindowTitle(c.title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}
