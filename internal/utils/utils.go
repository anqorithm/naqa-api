package utils

import "fmt"

// ###############################################################################
// Utility Functions
// ###############################################################################

func SafeString(v interface{}) string {
	if str, ok := v.(string); ok {
		return str
	}
	if num, ok := v.(float64); ok {
		return fmt.Sprintf("%.2f", num)
	}
	return ""
}
