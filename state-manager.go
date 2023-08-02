package gamecore

import (
	"bytes"
	"encoding/gob"
)

type StateManager interface {
	Get(key string, data any)
	Set(key string, data any)
	Destroy(key string)
}

type stateManager struct {
	control *control
	state   map[string][]byte
}

func createStateManager(control *control) *stateManager {
	return &stateManager{
		control: control,
		state:   make(map[string][]byte),
	}
}

func (m *stateManager) Get(key string, data any) {
	if _, ok := m.state[key]; !ok {
		return
	}
	stored := bytes.NewBuffer(m.state[key])
	decoder := gob.NewDecoder(stored)
	m.control.errorManager.Check(decoder.Decode(data))
}

func (m *stateManager) Set(key string, data any) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	m.control.errorManager.Check(encoder.Encode(data))
	m.state[key] = result.Bytes()
}

func (m *stateManager) Destroy(key string) {
	delete(m.state, key)
}
