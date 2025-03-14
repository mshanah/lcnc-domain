package domain

type Application struct {
	ID       string
	Name     string
	Type     string // e.g., Frontend, Backend, AI
	Language string // e.g., Go, Java, C++
}

func NewApplication(id, name, appType, language string) *Application {
	return &Application{
		ID:       id,
		Name:     name,
		Type:     appType,
		Language: language,
	}
}
