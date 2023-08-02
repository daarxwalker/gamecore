package gamecore

type StateManager interface {
}

type stateManager struct {
	control *control
}

func createStateManager(control *control) *stateManager {
	return &stateManager{
		control: control,
	}
}
