package handler_test

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/1995parham-teaching/k1s/internal/handler"
	"github.com/gofiber/fiber/v3"
)

func TestHealth_IsAlive(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	h := handler.NewHealth()
	h.Register(app.Group(""))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to test request: %v", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
	}
}

func TestHealth_Die(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	h := handler.NewHealth()
	h.Register(app.Group(""))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/die", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to test request: %v", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
	}

	if h.Status {
		t.Error("expected health status to be false after /die")
	}
}

func TestHealth_IsAlive_AfterDie(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	h := handler.NewHealth()
	h.Register(app.Group(""))

	// Set status to false to simulate dead state
	h.Status = false

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Use a short timeout since the handler sleeps for 1 minute when dead
	resp, err := app.Test(req, fiber.TestConfig{
		Timeout:       100 * time.Millisecond,
		FailOnTimeout: false,
	})
	if err == nil && resp != nil {
		defer func() { _ = resp.Body.Close() }()
		// If we got a response, it should be an error status
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", http.StatusInternalServerError, resp.StatusCode)
		}
	}
}

func TestHealth_TableDriven(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		route        string
		method       string
		expectedCode int
	}{
		{
			name:         "healthz endpoint returns 204",
			route:        "/healthz",
			method:       http.MethodGet,
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "die endpoint returns 204",
			route:        "/die",
			method:       http.MethodGet,
			expectedCode: http.StatusNoContent,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			app := fiber.New()
			h := handler.NewHealth()
			h.Register(app.Group(""))

			req, err := http.NewRequestWithContext(context.Background(), tc.method, tc.route, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("failed to test request: %v", err)
			}

			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != tc.expectedCode {
				body, _ := io.ReadAll(resp.Body)
				t.Errorf("expected status %d, got %d, body: %s", tc.expectedCode, resp.StatusCode, string(body))
			}
		})
	}
}
