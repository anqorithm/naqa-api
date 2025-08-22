package tests

import (
	"testing"

	"github.com/anqorithm/naqa-api/internal/utils"
)

func TestSafeString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "hello"},
		{123, "123"},
		{12.34, "12.34"},
		{int32(456), "456"},
		{int64(789), "789"},
		{nil, ""},
		{[]string{"test"}, ""},
	}

	for _, test := range tests {
		result := utils.SafeString(test.input)
		if result != test.expected {
			t.Errorf("SafeString(%v) = '%s', expected '%s'", test.input, result, test.expected)
		}
	}
}

func TestValidateRequest(t *testing.T) {
	type TestStruct struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
		Age   int    `validate:"min=18"`
	}

	validStruct := TestStruct{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   25,
	}

	errors := utils.ValidateRequest(validStruct)
	if errors != nil {
		t.Errorf("ValidateRequest() should return nil for valid struct, got %v", errors)
	}

	invalidStruct := TestStruct{
		Name:  "",
		Email: "invalid-email",
		Age:   15,
	}

	errors = utils.ValidateRequest(invalidStruct)
	if errors == nil {
		t.Error("ValidateRequest() should return errors for invalid struct")
	}

	if len(errors) != 3 {
		t.Errorf("Expected 3 validation errors, got %d", len(errors))
	}
}