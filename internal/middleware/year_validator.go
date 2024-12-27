package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var availableYears = []string{"2023", "2024"}

func ValidateYear() fiber.Handler {
	return func(c *fiber.Ctx) error {
		year := c.Params("year")
		if year == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Year parameter is required",
				"available_years": availableYears,
			})
		}

		valid := false
		for _, y := range availableYears {
			if y == year {
				valid = true
				break
			}
		}

		if !valid {
			return c.Status(400).JSON(fiber.Map{
				"error": fmt.Sprintf("Invalid year: %s", year),
				"message": "Please provide a valid year",
				"available_years": availableYears,
			})
		}

		return c.Next()
	}
}
