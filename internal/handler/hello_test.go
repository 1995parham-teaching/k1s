package handler_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/1995parham-teaching/k1s/internal/handler"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

const nameRaha = "Raha"

func TestHello_Say(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	logger := zap.NewNop()
	h := handler.NewHello("Hello", logger)
	h.Register(app.Group(""))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to test request: %v", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	bodyStr := string(body)

	if !strings.Contains(bodyStr, "Hello") {
		t.Errorf("expected body to contain 'Hello', got: %s", bodyStr)
	}

	if !strings.Contains(bodyStr, nameRaha) {
		t.Errorf("expected body to contain %q, got: %s", nameRaha, bodyStr)
	}
}

func TestHello_Say_CustomMessage(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	logger := zap.NewNop()
	h := handler.NewHello("Goodbye", logger)
	h.Register(app.Group(""))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to test request: %v", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	bodyStr := string(body)
	if !strings.Contains(bodyStr, "Goodbye") {
		t.Errorf("expected body to contain 'Goodbye', got: %s", bodyStr)
	}
}

func TestHello_TableDriven(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		greetingMessage string
		expectedCode    int
		expectedContain []string
	}{
		{"hello greeting", "Hello", http.StatusOK, []string{"Hello", nameRaha}},
		{"custom greeting", "Salam", http.StatusOK, []string{"Salam", nameRaha}},
		{"empty greeting", "", http.StatusOK, []string{nameRaha}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testHelloHandler(t, tc.greetingMessage, tc.expectedCode, tc.expectedContain)
		})
	}
}

func testHelloHandler(t *testing.T, msg string, expectedCode int, expectedContain []string) {
	t.Helper()

	app := fiber.New()
	logger := zap.NewNop()
	h := handler.NewHello(msg, logger)
	h.Register(app.Group(""))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed to test request: %v", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != expectedCode {
		t.Errorf("expected status %d, got %d", expectedCode, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	bodyStr := string(body)
	for _, expected := range expectedContain {
		if !strings.Contains(bodyStr, expected) {
			t.Errorf("expected body to contain '%s', got: %s", expected, bodyStr)
		}
	}
}
