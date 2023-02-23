package cmsapi

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"code"`
	Type       string `json:"type,omitempty"`
	Method     string `json:"method,omitempty"`
	Message    string `json:"message,omitempty"`
}

func NewAPIError(st int, tp string, mt string, mg error) *APIError {
	return &APIError{
		StatusCode: st,
		Type:       tp,
		Method:     mt,
		Message:    mg.Error(),
	}
}

func (a *APIError) Error() string {
	return fmt.Sprintf("status code: %v, error type: %v - method: %v - error message: %v", a.StatusCode, a.Type, a.Method, a.Message)
}

func InvalidJSON(mt string, err error) *APIError {
	return &APIError{
		StatusCode: http.StatusBadRequest,
		Type:       "invalid-json",
		Method:     mt,
		Message:    "Invalid or malformed JSON, " + err.Error(),
	}
}

func BadRequest(mt string, err error) *APIError {
	return &APIError{
		StatusCode: http.StatusBadRequest,
		Type:       "client-error",
		Method:     mt,
		Message:    "Cannot process current request, " + err.Error(),
	}
}

func InternalServerError(mt string, err error) *APIError {
	return &APIError{
		StatusCode: http.StatusInternalServerError,
		Type:       "server-error",
		Method:     mt,
		Message:    "Internal server error" + err.Error(),
	}
}
