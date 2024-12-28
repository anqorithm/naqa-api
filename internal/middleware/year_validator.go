package middleware

import (
	"fmt"

	"github.com/anqorithm/naqa-api/internal/constants"
	"github.com/anqorithm/naqa-api/internal/handlers"
	"github.com/anqorithm/naqa-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// ###############################################################################
// Year Validation Middleware
// ###############################################################################


func ValidateYear() fiber.Handler {
	return func(c *fiber.Ctx) error {
		year := c.Params("year")
		if year == "" {
			return handlers.SendError(c, fiber.StatusNotFound,
				models.ErrCodeValidationFailed,
				"Year parameter is required",
				fiber.Map{
					"available_years": constants.AvailableYears,
					"example":         fmt.Sprintf("/api/v1/stocks/year/%s", constants.AvailableYears[len(constants.AvailableYears)-1]),
				})
		}

		valid := false
		for _, y := range constants.AvailableYears {
			if y == year {
				valid = true
				break
			}
		}

		if !valid {
			return handlers.SendError(c, fiber.StatusNotFound,
				models.ErrCodeValidationFailed,
				fmt.Sprintf("Year '%s' not found", year),
				fiber.Map{
					"provided_year":   year,
					"available_years": constants.AvailableYears,
					"latest_year":     constants.AvailableYears[len(constants.AvailableYears)-1],
					"suggestion":      fmt.Sprintf("Try using the latest available year: %s", constants.AvailableYears[len(constants.AvailableYears)-1]),
				})
		}

		return c.Next()
	}
}
