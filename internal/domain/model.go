package domain

import "time"

// HealthStatus is returned by readiness checks.
type HealthStatus struct {
	Status  string    `json:"status"`
	Message string    `json:"message,omitempty"`
	Time    time.Time `json:"time"`
}

// ChatRequest is a minimal LLM gateway request shape.
type ChatRequest struct {
	Model    ModelID `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}