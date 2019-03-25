package shared

import (
	"fmt"
)

const (
	// ErrorDataNotFound error message when data doesn't exist
	ErrorDataNotFound = "data %s not found"
	// ErrorParameterInvalid error message for parameter is invalid
	ErrorParameterInvalid = "%s parameter is invalid"
	// ErrorParameterRequired error message for parameter is missing
	ErrorParameterRequired = "%s parameter is required"
	// ErrorParameterLength error message for parameter length is invalid
	ErrorParameterLength = "length of %s parameter exceeds the limit %d"
	// ErrorUnauthorized error message for unauthorized user
	ErrorUnauthorized = "you are not authorized"

	// ErrorRedisNil error for redis nil
	ErrorRedisNil = "redis: nil"
)

// ErrorAllowNumericOnly struct
type ErrorAllowNumericOnly struct {
	field   string
	message string
}

// ErrorValueShouldBool struct
type ErrorValueShouldBool struct {
	field   string
	message string
}

// NewErrorAllowNumericOnly ErrorRatingExceedLimit's constructor
func NewErrorAllowNumericOnly(field string) *ErrorAllowNumericOnly {
	return &ErrorAllowNumericOnly{field: field, message: "parameter %s should be number"}
}

// NewErrorValueShouldBool ErrorValueShouldBool's constructor
func NewErrorValueShouldBool(field string) *ErrorValueShouldBool {
	return &ErrorValueShouldBool{field: field, message: "%v value should be true or false"}
}

// Error function
func (e *ErrorAllowNumericOnly) Error() string {
	return fmt.Sprintf(e.message, e.field)
}

// Error function
func (e *ErrorValueShouldBool) Error() string {
	return fmt.Sprintf(e.message, e.field)
}
