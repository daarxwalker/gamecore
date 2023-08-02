package gamecore

type UpdateManager interface {
}

type updateManager struct {
	control *control
}

func createUpdateManager(control *control) *updateManager {
	return &updateManager{
		control: control,
	}
}
