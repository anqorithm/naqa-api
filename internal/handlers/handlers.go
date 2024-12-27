package handlers

import (
	"os"
	"time"

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
        "version":     "1.0.0",
        "env":         os.Getenv("APP_ENV"),
        "server_time": time.Now().Format(time.RFC3339),
        "request_id":  c.Get("X-Request-ID", uuid.New().String()),
        "endpoints": []string{
            "/api/v1/stocks",
            "/api/v1/health",
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

// HandleSomething handles an example endpoint
func (h *Handler) HandleSomething(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Handler example",
	})
}
