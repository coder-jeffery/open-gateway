package handler

import (
	"net/http"
	"time"

	"example.com/open-gateway/internal/domain"
	"example.com/open-gateway/internal/delivery/http/render"
)

// Health handles GET /healthz.
type Health struct{}

func (h Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	render.JSON(w, http.StatusOK, domain.HealthStatus{
		Status:  "ok",
		Message: "Hello World!",
		Time:    time.Now().UTC(),
	})
}