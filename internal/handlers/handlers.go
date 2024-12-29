package handlers

// ###############################################################################
// Handler Functions
// ###############################################################################

import (
	"fmt"
	"os"
	"time"

	"github.com/anqorithm/naqa-api/internal/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	db *mongo.Database
}

func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		db: db,
	}
}

// RootHandler handles the "/" endpoint
func (h *Handler) RootHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"app_name":      os.Getenv("APP_NAME"),
		"version":       os.Getenv("API_VERSION"),
		"description":   os.Getenv("APP_DESCRIPTION"),
		"request_time":  time.Now().Format(time.RFC3339),
		"environment":   os.Getenv("ENVIRONMENT"),
		"supported_api": []string{"/api/v1"},
	})
}

// ApiV1Handler handles the "/api/v1" endpoint
func (h *Handler) ApiV1Handler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":      "active",
		"message":     "Welcome to NAQA API v1",
		"version":     os.Getenv("API_VERSION"),
		"env":         os.Getenv("ENVIRONMENT"),
		"server_time": time.Now().Format(time.RFC3339),
		"request_id":  c.Get("X-Request-ID", uuid.New().String()),
		"endpoints": []string{
			"/api/v1/stocks",
			"/api/v1/health",
			"/api/v1/metrics",
		},
	})
}

// HealthCheckHandler handles the health check endpoint
func (h *Handler) HealthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "OK",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func (h *Handler) GetStocksBaseHandler(c *fiber.Ctx) error {
	latestYear := constants.AvailableYears[len(constants.AvailableYears)-1]
	return c.JSON(fiber.Map{
		"status":          "success",
		"message":         "Welcome to Stocks API",
		"available_years": constants.AvailableYears,
		"endpoints": []map[string]interface{}{
			{
				"name":        "Get Stocks",
				"path":        "/api/v1/stocks/year/{year}",
				"method":      "GET",
				"description": "Get all stocks for a specific year",
				"example":     fmt.Sprintf("/api/v1/stocks/year/%s", latestYear),
			},
			{
				"name":        "Search Stocks",
				"path":        "/api/v1/stocks/year/{year}/search",
				"method":      "GET",
				"description": "Search stocks with filters",
				"example":     fmt.Sprintf("/api/v1/stocks/year/%s/search?sector=الطاقة&sharia_opinion=نقية", latestYear),
				"parameters": []string{
					"name", "code", "sector", "sharia_opinion",
				},
			},
			{
				"name":        "Calculate Purification",
				"path":        "/api/v1/stocks/year/{year}/calculate-purification",
				"method":      "POST",
				"description": "Calculate stock purification amount",
				"example": map[string]interface{}{
					"url": fmt.Sprintf("/api/v1/stocks/year/%s/calculate-purification", latestYear),
					"body": map[string]interface{}{
						"start_date":       "2023-01-01",
						"end_date":         "2023-12-31",
						"number_of_stocks": 100,
						"stock_code":       "1111",
					},
				},
			},
		},
		"documentation": "https://github.com/anqorithm/naqa-api",
	})
}
