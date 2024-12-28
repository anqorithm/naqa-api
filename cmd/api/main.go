package main

import (
	"log"
	"os"
	"strings"

	"github.com/anqorithm/naqa-api/internal/config"
	"github.com/anqorithm/naqa-api/internal/middleware"
	"github.com/anqorithm/naqa-api/internal/routes"
	"github.com/anqorithm/naqa-api/internal/seeders"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// ###############################################################################
	// Environment and Configuration Setup
	// ###############################################################################
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
	mongoConfig := config.NewMongoConfig()
	db, err := config.ConnectDB(mongoConfig)
	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}

	// ###############################################################################
	// Data Seeding
	// ###############################################################################
	shouldSeed := strings.ToLower(os.Getenv("SEED_DATA")) == "true"
	if shouldSeed {
		log.Println("Starting data seeding process...")
		if err := seeders.LoadDataSources(db); err != nil {
			log.Printf("Warning: Error seeding data: %v", err)
		}
	} else {
		log.Println("Skipping data seeding (SEED_DATA is not set to 'true')")
	}

	// ###############################################################################
	// Application Setup and Middleware Configuration
	// ###############################################################################
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(middleware.Logger())
	app.Use(middleware.Security())
	app.Use(middleware.Compress())
	app.Use(middleware.CORS())
	app.Use(middleware.RateLimit())
	app.Use("/api/v1/metrics", middleware.Monitor())

	// ###############################################################################
	// Router Setup and Server Startup
	// ###############################################################################
	router := routes.NewRouter(app, db)
	router.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
