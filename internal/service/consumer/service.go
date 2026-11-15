package consumer

import "example.com/open-gateway/internal/domain"

// Service manages tenants, virtual keys, and quotas.
type Service struct{}

func New() *Service {
	return &Service{}
}

// Authenticate validates a virtual API key (stub).
func (s *Service) Authenticate(_ domain.APIKey) (domain.TenantID, error) {
	return "", domain.ErrUnauthorized
}