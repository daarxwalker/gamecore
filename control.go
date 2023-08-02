package gamecore

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Control interface {
	Window() WindowManager
}

type control struct {
	screen        *ebiten.Image
	gamecore      *gamecore
	errorManager  *errorManager
	renderManager *renderManager
	stateManager  *stateManager
	updateManager *updateManager
	windowManager *windowManager
}

func createControl(gamecore *gamecore) *control {
	c := &control{gamecore: gamecore}
	c.errorManager = createErrorManager(c)
	c.renderManager = createRenderManager(c)
	c.updateManager = createUpdateManager(c)
	c.stateManager = createStateManager(c)
	return c
}

func (c *control) Window() WindowManager {
	return c.windowManager
}
