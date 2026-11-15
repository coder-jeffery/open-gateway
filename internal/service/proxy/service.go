package proxy

// Service handles generic upstream proxying, streaming, and header transforms.
type Service struct{}

func New() *Service {
	return &Service{}
}