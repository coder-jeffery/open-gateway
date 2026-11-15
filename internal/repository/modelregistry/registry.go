package modelregistry

// Entry describes upstream model metadata and provider credentials reference.
type Entry struct {
	ID       string
	Provider string
	Model    string
}

// Registry holds model metadata (stub).
type Registry struct{}

func New() *Registry {
	return &Registry{}
}