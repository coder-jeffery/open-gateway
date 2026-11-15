package llm

// Service implements LLM gateway: routing, fallback, token accounting, budget.
type Service struct{}

func New() *Service {
	return &Service{}
}