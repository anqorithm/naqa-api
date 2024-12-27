package middleware

import (
	"fmt"

	"github.com/anqorithm/naqa-api/internal/handlers"
	"github.com/anqorithm/naqa-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// ###############################################################################
// Year Validation Middleware
// ###############################################################################

var availableYears = []string{"2018", "2019", "2020", "2021", "2022", "2023"}

func ValidateYear() fiber.Handler {
	return func(c *fiber.Ctx) error {
		year := c.Params("year")
		if year == "" {
			return handlers.SendError(c, fiber.StatusBadRequest, 
				models.ErrCodeValidationFailed,
				"Year parameter is required",
				fiber.Map{
					"available_years": availableYears,
					"example": fmt.Sprintf("/api/v1/stocks/year/%s", availableYears[len(availableYears)-1]),
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
			return handlers.SendError(c, fiber.StatusBadRequest, 
				models.ErrCodeValidationFailed,
				fmt.Sprintf("The year '%s' is not supported", year),
				fiber.Map{
					"provided_year":    year,
					"available_years": availableYears,
					"latest_year":      availableYears[len(availableYears)-1],
					"suggestion":       fmt.Sprintf("Try using the latest available year: %s", availableYears[len(availableYears)-1]),
				})
		}

		return c.Next()
	}
}
