package middleware

// ###############################################################################
// Middleware Functions
// ###############################################################################

import (
	"fmt"
	"time"

	"github.com/anqorithm/naqa-api/internal/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Logger returns a logger middleware
func Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	})
}

// Security returns helmet middleware
func Security() fiber.Handler {
	return helmet.New()
}

// Compress returns compression middleware
func Compress() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
}

// CORS returns CORS middleware
func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	})
}

// RateLimit returns rate limiter middleware
func RateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
	})
}

// Monitor returns a monitor middleware
func Monitor() fiber.Handler {
    return monitor.New(monitor.Config{
        Title:      fmt.Sprintf("%s | %s", constants.MonitorTitleAr, constants.MonitorTitleEn),
        Refresh:    3 * time.Second,
        Next:       nil,
        APIOnly:    false,
    })
}