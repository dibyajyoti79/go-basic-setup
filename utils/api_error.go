package utils

import (
	"fmt"
	"runtime"
)

type ApiError struct {
	StatusCode int      `json:"statusCode"`
	Success    bool     `json:"success"`
	Message    string   `json:"message"`
	Errors     []string `json:"errors"`
	Stack      string   `json:"stack,omitempty"`
}

func NewApiError(statusCode int, message string, errors []string) *ApiError {
	if message == "" {
		message = "Something went wrong"
	}
	_, file, line, _ := runtime.Caller(1)
	stack := fmt.Sprintf("%s:%d", file, line)
	return &ApiError{
		StatusCode: statusCode,
		Success:    false,
		Message:    message,
		Errors:     errors,
		Stack:      stack,
	}
}

// Implement the error interface
func (e *ApiError) Error() string {
	return e.Message
}
