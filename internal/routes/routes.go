package routes

import (
	"github.com/anqorithm/naqa-api/internal/handlers"
	"github.com/anqorithm/naqa-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	app *fiber.App
	h   *handlers.Handler
}

func NewRouter(app *fiber.App, db *mongo.Database) *Router {
	return &Router{
		app: app,
		h:   handlers.NewHandler(db),
	}
}

func (r *Router) SetupRoutes() {
	// Global middleware
	r.app.Use(middleware.Logger())

	// Root route
	r.app.Get("/", r.h.RootHandler)

	// API v1 routes
	v1 := r.app.Group("/api/v1")
	r.setupV1Routes(v1)

	// API routes
	api := r.app.Group("/api")

	// Health check route
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}

func (r *Router) setupV1Routes(v1 fiber.Router) {
	// Base API info
	v1.Get("/", r.h.ApiV1Handler)

	// Health check
	v1.Get("/health", r.h.HealthCheckHandler)

	// Resources
	r.setupUserRoutes(v1)
	r.setupTaskRoutes(v1)

	// Stock market routes
	r.setupStockRoutes(v1)
}

func (r *Router) setupUserRoutes(v1 fiber.Router) {
	users := v1.Group("/users")
	users.Get("/", r.h.GetUsersHandler)
}

func (r *Router) setupTaskRoutes(v1 fiber.Router) {
	tasks := v1.Group("/tasks")
	tasks.Get("/", r.h.GetTasksHandler)
}

func (r *Router) setupStockRoutes(v1 fiber.Router) {
	stocks := v1.Group("/stocks")
	
	// Get all stocks for a specific year
	stocks.Get("/year/:year", middleware.ValidateYear(), r.h.GetStocksByYearHandler)
	
	// Search stocks by various parameters
	stocks.Get("/year/:year/search", middleware.ValidateYear(), r.h.SearchStocksHandler)
}