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
    // ###############################################################################
    // Environment and Configuration Setup
    // ###############################################################################
    if os.Getenv("ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Printf("Warning: Error loading .env file: %v", err)
        }
    }

    mongoConfig := config.NewMongoConfig()
    db, err := config.ConnectDB(mongoConfig)
    if err != nil {
        log.Fatal("Database Connection Error: ", err)
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
    app.Use("/metrics", middleware.Monitor())

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