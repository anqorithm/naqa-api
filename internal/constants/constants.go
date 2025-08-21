package constants

// ###############################################################################
// Constants
// ###############################################################################

// Date format for validation and parsing
const DateFormat = "2006-01-02"

// Error codes
const (
	ErrCodeInvalidRequest   = "INVALID_REQUEST"
	ErrCodeValidationFailed = "VALIDATION_FAILED"
	ErrCodeDatabaseError    = "DATABASE_ERROR"
	ErrCodeNotFound         = "NOT_FOUND"
	ErrCodeInvalidData      = "INVALID_DATA"
)

// Error messages
const (
	MsgInvalidRequestBody      = "Invalid request body"
	MsgValidationFailed        = "Validation failed"
	MsgInvalidStartDateFormat  = "Invalid start date format. Use YYYY-MM-DD."
	MsgInvalidEndDateFormat    = "Invalid end date format. Use YYYY-MM-DD."
	MsgEndDateAfterStartDate   = "End date must be greater than start date."
	MsgStockNotFound           = "Stock not found"
	MsgInvalidPurificationRate = "Invalid purification rate in database"
)

// Available years for stock data
var AvailableYears = []string{"2015", "2016", "2017", "2018", "2019", "2020", "2021", "2022", "2023", "2024"}

// Monitor titles
const (
	MonitorTitleAr = "لوحة مراقبة نقاء"
	MonitorTitleEn = "NAQA Monitoring Dashboard"
)

// Error Messages
const (
	ErrDatabaseConnection = "Database Connection Error"
	ErrDataSeeding        = "Error seeding data"
	ErrInternalServer     = "Internal Server Error"
	ErrInvalidInput       = "Invalid input data"
	ErrNotFound           = "Resource not found"
	ErrUnauthorized       = "Unauthorized access"
)

// Success Messages
const (
	SuccessDataSeeded    = "Data seeding completed successfully"
	SuccessServerStarted = "Server started successfully"
	SuccessHealthCheck   = "Health check passed"
)

// Info Messages
const (
	InfoSkippingSeeding = "Skipping data seeding (SEED_DATA is not set to 'true')"
	InfoStartingSeeding = "Starting data seeding process..."
	InfoServerStarting  = "Server starting on port %s"
)

// Validation Messages
const (
	ValidateRequiredField = "This field is required"
	ValidateInvalidEmail  = "Invalid email format"
	ValidateInvalidInput  = "Invalid input format"
)