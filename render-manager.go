package gamecore

import (
	"fmt"
	
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type RenderManager interface {
	Debug(value any)
}

type renderManager struct {
	control *control
}

func createRenderManager(control *control) *renderManager {
	return &renderManager{
		control: control,
	}
}

func (m *renderManager) Debug(value any) {
	ebitenutil.DebugPrint(m.control.screen, fmt.Sprintf("%v", value))
}
