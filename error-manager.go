package gamecore

type ErrorManager interface {
	Check(err error)
}

type errorManager struct {
	control *control
}

func createErrorManager(control *control) *errorManager {
	return &errorManager{
		control: control,
	}
}

func (em *errorManager) Check(err error) {
	if err == nil {
		return
	}
	panic(err)
}
