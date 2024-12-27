package main

import (
	"log"
	"os"

	"github.com/anqorithm/naqa-api/internal/config"
	"github.com/anqorithm/naqa-api/internal/middleware"
	"github.com/anqorithm/naqa-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Initialize MongoDB
	mongoConfig := config.NewMongoConfig()
	db, err := config.ConnectDB(mongoConfig)
	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Global Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Security())
	app.Use(middleware.Compress())
	app.Use(middleware.CORS())
	app.Use(middleware.RateLimit())	

	// Monitor dashboard at /metrics
	app.Use("/metrics", middleware.Monitor())

	// Setup routes
	router := routes.NewRouter(app, db)
	router.SetupRoutes()

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
