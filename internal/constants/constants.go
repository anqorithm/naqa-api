package constants

// ###############################################################################
// Constants
// ###############################################################################

// Date format for validation and parsing
const DateFormat = "2006-01-02"

// Error codes
const (
	ErrCodeInvalidRequest      = "INVALID_REQUEST"
	ErrCodeValidationFailed    = "VALIDATION_FAILED"
	ErrCodeDatabaseError       = "DATABASE_ERROR"
	ErrCodeNotFound            = "NOT_FOUND"
	ErrCodeInvalidData         = "INVALID_DATA"
)

// Error messages
const (
	MsgInvalidRequestBody        = "Invalid request body"
	MsgValidationFailed          = "Validation failed"
	MsgInvalidStartDateFormat    = "Invalid start date format. Use YYYY-MM-DD."
	MsgInvalidEndDateFormat      = "Invalid end date format. Use YYYY-MM-DD."
	MsgEndDateAfterStartDate     = "End date must be greater than start date."
	MsgStockNotFound             = "Stock not found"
	MsgInvalidPurificationRate   = "Invalid purification rate in database"
)

// Available years for stock data
var AvailableYears = []string{"2021", "2022", "2023", "2024", "2025"}