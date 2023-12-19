package utils

import (
	"encoding/json"
	"fmt"
)

// ResponseError format
type ResponseError struct {
	Err *ErrorBody `json:"error"`
}

// ErrorBody format to return as response
type ErrorBody struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

// ErrorCode for all services
type ErrorCode map[string]map[string]map[string]int

// AppError standardize error object
type AppError struct {
	// Code can be used as HTTP status
	Code int
	// Unique error code to identify service, module, and category of error
	ErrorCode string
	// Original error object
	Err error
	// [Optional] error message, if omitted will be taken from Err
	Message string
	// [Optional] additional error details
	Details []string
}

func (err *AppError) Error() string {
	return fmt.Sprintf("Error: %s %s %v", err.ErrorCode, err.Message, err.Details)
}

// ToJSON format error as json response
func (err *AppError) ToJSON() []byte {
	if err.Message == "" {
		err.Message = err.Err.Error()
	}

	eb := &ErrorBody{
		Code:    err.ErrorCode,
		Message: err.Message,
		Details: err.Details,
	}

	respErr := &ResponseError{
		Err: eb,
	}

	r, je := json.Marshal(respErr)
	if je != nil {
		return []byte(err.Error())
	}

	return r
}
