package kv

// Store wraps Redis for rate limits, cache, and token counters.
type Store struct {
	Addr string
}

func New(addr string) *Store {
	return &Store{Addr: addr}
}