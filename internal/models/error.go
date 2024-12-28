package models

// ###############################################################################
// Error Response Model
// ###############################################################################

type ErrorResponse struct {
	Status  string      `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Common error codes
const (
	ErrCodeInvalidRequest    = "INVALID_REQUEST"
	ErrCodeNotFound          = "NOT_FOUND"
	ErrCodeInternalError     = "INTERNAL_ERROR"
	ErrCodeValidationFailed  = "VALIDATION_FAILED"
	ErrCodeDatabaseError     = "DATABASE_ERROR"
	ErrCodeInvalidDateFormat = "INVALID_DATE_FORMAT"
	ErrCodeInvalidData       = "INVALID_DATA"
)