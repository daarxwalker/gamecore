package gamecore

type WindowManager interface {
}

type windowManager struct {
	width  int
	height int
}

func createWindowManager(width, height int) *windowManager {
	return &windowManager{
		width:  width,
		height: height,
	}
}
