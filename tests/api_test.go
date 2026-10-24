package tests

import (
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/anqorithm/naqa-api/internal/routes"
	"github.com/anqorithm/naqa-api/internal/seeders"
	"github.com/gofiber/fiber/v3"
)

// newTestApp seeds a throwaway database and returns the app with all routes, or skips without MongoDB
func newTestApp(t *testing.T) *fiber.App {
	db := setupTestDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := db.Client().Ping(ctx, nil); err != nil {
		t.Skip("MongoDB not available for testing")
	}
	t.Cleanup(func() { db.Drop(context.Background()) })

	wd, _ := os.Getwd()
	if err := os.Chdir(".."); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(wd)
	if err := seeders.LoadDataSources(db); err != nil {
		t.Fatal(err)
	}

	app := fiber.New()
	routes.NewRouter(app, db).SetupRoutes()
	return app
}

func call(t *testing.T, app *fiber.App, method, path, body string) (int, map[string]interface{}) {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, fiber.TestConfig{Timeout: 10 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	raw, _ := io.ReadAll(resp.Body)
	var out map[string]interface{}
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatalf("%s %s: invalid JSON %q", method, path, raw)
	}
	return resp.StatusCode, out
}

func TestAPIEndpoints(t *testing.T) {
	app := newTestApp(t)

	if code, body := call(t, app, "GET", "/api/v1/health", ""); code != 200 || body["status"] != "OK" {
		t.Fatalf("health: %d %v", code, body)
	}

	code, body := call(t, app, "GET", "/api/v1/stocks", "")
	if years, _ := body["available_years"].([]interface{}); code != 200 || len(years) == 0 || years[len(years)-1] != "2025" {
		t.Fatalf("stocks base: %d %v", code, body["available_years"])
	}

	code, body = call(t, app, "GET", "/api/v1/stocks/year/2025", "")
	if stocks, _ := body["stocks"].([]interface{}); code != 200 || len(stocks) != 382 {
		t.Fatalf("2025 stocks: %d, %d rows", code, len(stocks))
	}

	code, body = call(t, app, "GET", "/api/v1/stocks/year/2025/search?code=2222", "")
	stocks, _ := body["stocks"].([]interface{})
	if code != 200 || len(stocks) != 1 {
		t.Fatalf("search 2222: %d %v", code, body)
	}
	aramco := stocks[0].(map[string]interface{})
	if aramco["name_en"] != "Saudi Arabian Oil Co." || aramco["purification"] != "0.05" || aramco["_id"] == "" {
		t.Fatalf("aramco: %v", aramco)
	}

	if code, body = call(t, app, "GET", "/api/v1/stocks/year/1999", ""); code != 404 || body["code"] != "VALIDATION_FAILED" {
		t.Fatalf("invalid year: %d %v", code, body)
	}

	// 100 shares of Aramco held for all of 2025 at 0.05 per share
	code, body = call(t, app, "POST", "/api/v1/stocks/year/2025/calculate-purification",
		`{"start_date":"2025-01-01","end_date":"2025-12-31","number_of_stocks":100,"stock_code":"2222"}`)
	if code != 200 || body["purification_amount"].(float64) < 4.99 || body["purification_amount"].(float64) > 5.01 || body["days_held"].(float64) != 365 {
		t.Fatalf("purification: %d %v", code, body)
	}

	if code, body = call(t, app, "POST", "/api/v1/stocks/year/2025/calculate-purification", `{"stock_code":"2222"}`); code != 400 || body["code"] != "VALIDATION_FAILED" {
		t.Fatalf("validation: %d %v", code, body)
	}
	if code, _ = call(t, app, "POST", "/api/v1/stocks/year/2025/calculate-purification", `not json`); code != 400 {
		t.Fatalf("bad body: %d", code)
	}
}
