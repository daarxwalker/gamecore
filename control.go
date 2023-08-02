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
	fontManager   *fontManager
	renderManager *renderManager
	stateManager  *stateManager
	updateManager *updateManager
	windowManager *windowManager
	textRenderer  *textRenderer
}

func createControl(gamecore *gamecore) *control {
	c := &control{gamecore: gamecore}
	c.errorManager = createErrorManager(c)
	c.fontManager = createFontManager(c)
	c.renderManager = createRenderManager(c)
	c.updateManager = createUpdateManager(c)
	c.stateManager = createStateManager(c)
	c.textRenderer = createTextRenderer(c)
	return c
}

func (c *control) Window() WindowManager {
	return c.windowManager
}
