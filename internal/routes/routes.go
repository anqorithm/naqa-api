package routes

// ###############################################################################
// Routes
// ###############################################################################

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

// ###############################################################################
// Router Setup and Server Startup
// ###############################################################################
func (r *Router) SetupRoutes() {
	// Root route
	r.app.Get("/", r.h.RootHandler)

	// API Version 1 Routes
	v1 := r.app.Group("/api/v1")
	v1.Get("/", r.h.ApiV1Handler)
	v1.Get("/health", r.h.HealthCheckHandler)
	r.setupV1Routes(v1)
}

// ###############################################################################
// API Version 1 Route Setup
// ###############################################################################
func (r *Router) setupV1Routes(v1 fiber.Router) {
	r.setupStockRoutes(v1)
}

// ###############################################################################
// Stock Routes Setup
// ###############################################################################
func (r *Router) setupStockRoutes(v1 fiber.Router) {
	stocks := v1.Group("/stocks")

	stocks.Get("/", r.h.GetStocksBaseHandler)

	// Year-Based Stock Routes
	yearGroup := stocks.Group("/year/:year", middleware.ValidateYear())

	// Get all stocks for a specific year
	yearGroup.Get("/", r.h.GetStocksByYearHandler)

	// Search stocks by various parameters
	yearGroup.Get("/search", r.h.SearchStocksHandler)

	// Calculate purification for a specific stock
	yearGroup.Post("/calculate-purification", r.h.CalculatePurificationHandler)
}

