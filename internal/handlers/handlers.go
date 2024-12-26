package handlers

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
		"message": "Welcome to API v1",
		"endpoints": []string{
			"/api/v1/health",
			"/api/v1/users",
			"/api/v1/tasks",
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

// GetUsersHandler handles the /api/v1/users endpoint
func (h *Handler) GetUsersHandler(c *fiber.Ctx) error {
	collection := h.db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching users",
		})
	}
	defer cursor.Close(ctx)

	var users []fiber.Map
	if err := cursor.All(ctx, &users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error parsing users data",
		})
	}

	return c.JSON(users)
}

// GetTasksHandler handles the /api/v1/tasks endpoint
func (h *Handler) GetTasksHandler(c *fiber.Ctx) error {
	collection := h.db.Collection("tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching tasks",
		})
	}
	defer cursor.Close(ctx)

	var tasks []fiber.Map
	if err := cursor.All(ctx, &tasks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error parsing tasks data",
		})
	}

	return c.JSON(tasks)
}
