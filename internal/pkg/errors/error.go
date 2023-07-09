package errors

import (
	"fmt"
	"strings"
)

// ErrorType is the type of error
type ErrorType string

const (
	NotFound      ErrorType = "NotFound"
	InternalError ErrorType = "InternalError"
)

type Error interface {
	Type() ErrorType
	Error() string
}

type customError struct {
	errorType ErrorType
	causes    []Cause
}

type Cause struct {
	Code    ErrorCode
	Message string
}

// New creates a new customError
func (errorType ErrorType) NewErrorF(causes ...Cause) Error {
	return customError{errorType: errorType, causes: causes}
}

func NewError() Error {
	return customError{
		causes: []Cause{},
	}
}

// Creates a new error cause with formatted message and code
func NewErrorCausef(code ErrorCode, format string, args ...interface{}) Cause {
	return Cause{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

func (err customError) Error() string {
	var causes []string
	for _, ec := range err.causes {
		causes = append(causes, ec.Message)
	}
	return strings.Join(causes, "; ")
}

// GetType returns the error type
func (err customError) Type() ErrorType {
	return err.errorType
}
