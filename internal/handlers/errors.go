package handlers

// ###############################################################################
// Error Handling Functions
// ###############################################################################

import (
	"github.com/anqorithm/naqa-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// SendError sends an error response to the client
func SendError(c *fiber.Ctx, status int, code string, message string, details interface{}) error {
	return c.Status(status).JSON(&models.ErrorResponse{
		Status:  "error",
		Code:    code,
		Message: message,
		Details: details,
	})
}

// Common error codes
const (
	ErrCodeInvalidRequest     = "INVALID_REQUEST"
	ErrCodeNotFound          = "NOT_FOUND"
	ErrCodeInternalError     = "INTERNAL_ERROR"
	ErrCodeValidationFailed  = "VALIDATION_FAILED"
	ErrCodeDatabaseError     = "DATABASE_ERROR"
	ErrCodeInvalidDateFormat = "INVALID_DATE_FORMAT"
)
