package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status  string      `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func NewErrorResponse(code string, message string, details interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:  "error",
		Code:    code,
		Message: message,
		Details: details,
	}
}

func SendError(c *fiber.Ctx, status int, code string, message string, details interface{}) error {
	return c.Status(status).JSON(NewErrorResponse(code, message, details))
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
