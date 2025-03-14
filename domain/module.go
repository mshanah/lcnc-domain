package domain

type Module struct {
	ID           string
	Name         string
	Applications []*Application
}

func NewModule(id, name string) *Module {
	return &Module{
		ID:           id,
		Name:         name,
		Applications: []*Application{},
	}
}

func (m *Module) AddApplication(app *Application) {
	m.Applications = append(m.Applications, app)
}
