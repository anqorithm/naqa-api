package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/anqorithm/naqa-api/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func TestRateLimit(t *testing.T) {
	app := fiber.New()
	app.Use(middleware.RateLimit())
	ok := func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) }
	app.Get("/api/v1/stocks", ok)
	app.Get("/api/v1/health", ok)

	status := func(path string) int {
		resp, err := app.Test(httptest.NewRequest("GET", path, nil))
		if err != nil {
			t.Fatal(err)
		}
		return resp.StatusCode
	}

	for i := 0; i < 100; i++ {
		if got := status("/api/v1/stocks"); got != fiber.StatusOK {
			t.Fatalf("request %d: status = %d, want 200", i+1, got)
		}
	}
	if got := status("/api/v1/stocks"); got != fiber.StatusTooManyRequests {
		t.Fatalf("request 101: status = %d, want 429", got)
	}
	if got := status("/api/v1/health"); got != fiber.StatusOK {
		t.Fatalf("health after limit: status = %d, want 200", got)
	}
}
