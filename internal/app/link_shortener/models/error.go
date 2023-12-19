package models

import (
	"errors"
	"fmt"
	"net/http"

	"link-shortener/internal/pkg/utils"
)

var errorCode = &utils.ErrorCode{
	"General": {
		"InternalServerError": {
			"module":   0,
			"category": 0,
			"code":     http.StatusInternalServerError,
		},
		"JSONError": {
			"module":   0,
			"category": 1,
			"code":     http.StatusInternalServerError,
		},
	},
	"Link": {
		"InternalServerError": {
			"module":   1,
			"category": 0,
			"code":     http.StatusInternalServerError,
		},
		"NotFound": {
			"module":   1,
			"category": 1,
			"code":     http.StatusNotFound,
		},
		"InvalidParameters": {
			"module":   1,
			"category": 2,
			"code":     http.StatusBadRequest,
		},
	},
	"User": {
		"InternalServerError": {
			"module":   2,
			"category": 0,
			"code":     http.StatusInternalServerError,
		},
		"InvalidParameters": {
			"module":   2,
			"category": 1,
			"code":     http.StatusBadRequest,
		},
	},
}

// WrapError turn ErrorCode to AppError
func WrapError(module string, category string, err error, details []string) *utils.AppError {
	if aerr, ok := err.(*utils.AppError); ok {
		return aerr
	}

	errModule := (*errorCode)[module]
	if errModule == nil {
		module = "General"
	}
	e := errModule[category]
	if e == nil {
		category = "InternalServerError"
	}

	if err == nil {
		err = errors.New(category)
	}

	code := fmt.Sprintf("gfi%d%s%s", e["code"], fmt.Sprintf("%02d", e["module"]), fmt.Sprintf("%02d", e["category"]))

	return &utils.AppError{
		Code:      e["code"],
		ErrorCode: code,
		Message:   err.Error(),
		Err:       err,
		Details:   details,
	}
}
