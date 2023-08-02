package gamecore

type UpdateManager interface {
	State() StateManager
}

type updateManager struct {
	control *control
}

func createUpdateManager(control *control) *updateManager {
	return &updateManager{
		control: control,
	}
}

func (m *updateManager) State() StateManager {
	return m.control.stateManager
}
