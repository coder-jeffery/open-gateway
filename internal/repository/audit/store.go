package audit

import "time"

// Record is a prompt / call audit entry.
type Record struct {
	TenantID  string
	Model     string
	CreatedAt time.Time
}

// Store persists audit records (stub).
type Store struct{}

func New() *Store {
	return &Store{}
}