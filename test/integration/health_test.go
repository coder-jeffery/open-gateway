//go:build integration

package integration

import (
	"net/http"
	"testing"
	"time"
)

func TestHealthz(t *testing.T) {
	base := "http://127.0.0.1:8080"
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(base + "/healthz")
	if err != nil {
		t.Skipf("server not running: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
}