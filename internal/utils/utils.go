package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func SafeString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%.2f", val)
	case int, int32, int64:
		return fmt.Sprintf("%d", val)
	default:
		return ""
	}
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateRequest(req interface{}) map[string]string {
    err := validate.Struct(req)
    if err == nil {
        return nil
    }
    validationErrors := make(map[string]string)
    for _, validationErr := range err.(validator.ValidationErrors) {
        fmt.Printf("Validation Error: Field=%s, Tag=%s, Param=%s\n",
            validationErr.Field(), validationErr.Tag(), validationErr.Param())
        validationErrors[validationErr.Field()] = fmt.Sprintf("%s validation failed on '%s'", validationErr.Field(), validationErr.Tag())
    }
    return validationErrors
}

