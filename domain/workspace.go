package domain

type Workspace struct {
	ID      string
	Name    string
	Modules []*Module
}

func NewWorkspace(id, name string) *Workspace {
	return &Workspace{ID: id, Name: name, Modules: []*Module{}}
}

func (w *Workspace) AddModule(module *Module) {
	w.Modules = append(w.Modules, module)
}
