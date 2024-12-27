package main

import (
	"log"
	"os"

	"github.com/anqorithm/naqa-api/internal/config"
	"github.com/anqorithm/naqa-api/internal/middleware"
	"github.com/anqorithm/naqa-api/internal/routes"

	_ "github.com/anqorithm/naqa-api/docs" // import swagger docs
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// @title           Naqa API
// @version         1.0
// @description     Stock Market API Service
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1
// @schemes   http https
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
