package gamecore

type Component interface {
	Init()
	Update(m UpdateManager)
	Render(m RenderManager)
}
