package gamecore

type Module interface {
	Component(component Component) Module
	Name(name string) Module
	
	getPtr() *module
}

type module struct {
	control    *control
	components []Component
	name       string
}

func createModule(control *control) *module {
	m := &module{
		control:    control,
		components: make([]Component, 0),
	}
	return m
}

func (m *module) Component(component Component) Module {
	m.components = append(m.components, component)
	return m
}

func (m *module) Name(name string) Module {
	m.name = name
	return m
}

func (m *module) getPtr() *module {
	return m
}

func (m *module) initComponents() {
	for _, c := range m.components {
		c.Init()
	}
}

func (m *module) updateComponents() {
	for _, c := range m.components {
		c.Update(m.control.updateManager)
	}
}

func (m *module) renderComponents() {
	for _, c := range m.components {
		c.Render(m.control.renderManager)
	}
}
