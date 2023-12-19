package network

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	utils "link-shortener/internal/pkg/utils"
)

// RequestErrorCode for request module
const RequestErrorCode = "gfx5002000"

// Request call to other service
type Request struct {
	Ctx     context.Context
	Name    string
	Method  string
	URL     string
	Path    string
	Headers map[string]string
	Data    interface{}
	Timeout int
}

// ServiceRequest call other service via HTTP
func (request Request) ServiceRequest() ([]byte, *utils.AppError) {
	var payload io.Reader

	var respErr *utils.ResponseError = &utils.ResponseError{}
	var appErr *utils.AppError = &utils.AppError{
		Code:      500,
		ErrorCode: RequestErrorCode,
	}

	if request.Method == "GET" {
		payload = nil
	} else {
		requestBody, err := json.Marshal(request.Data)

		if err != nil {
			appErr.Err = err
			return nil, appErr
		}

		payload = bytes.NewBuffer(requestBody)
	}

	var timeout time.Duration
	if request.Timeout > 0 {
		timeout = time.Duration(request.Timeout) * time.Second
	} else {
		timeout = time.Duration(60 * time.Second)
	}

	client := http.Client{
		Timeout: timeout,
	}

	reqID := utils.RandString(8, true)

	utils.Log(request.Ctx, "Request %s [%s] %s %s, %s", request.Name, reqID, request.Method, request.Path, payload)

	httpRequest, err := http.NewRequest(request.Method, request.URL+request.Path, payload)
	httpRequest.Header.Set("Content-Type", "application/json")
	for key, value := range request.Headers {
		httpRequest.Header.Set(key, value)
	}
	httpRequest.Header.Set("X-Request-Id", reqID)

	if err != nil {
		appErr.Err = err
		return nil, appErr
	}

	resp, err := client.Do(httpRequest)
	if err != nil {
		appErr.Err = err
		return nil, appErr
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		appErr.Err = err
		return nil, appErr
	}

	utils.Log(request.Ctx, "Response %s [%s] %s", request.Name, reqID, string(body))

	if resp.StatusCode >= http.StatusBadRequest {
		err := json.Unmarshal(body, &respErr)
		if err != nil {
			appErr.Err = err
			return nil, appErr
		}

		appErr.Code = resp.StatusCode
		appErr.ErrorCode = respErr.Err.Code
		appErr.Message = respErr.Err.Message
		appErr.Details = respErr.Err.Details

		return nil, appErr
	}

	return body, nil
}
