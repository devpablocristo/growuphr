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

func NewAPIError(st int, tp string, mt string, mg string) *APIError {
	return &APIError{
		StatusCode: st,
		Type:       tp,
		Method:     mt,
		Message:    mg,
	}
}

func (a *APIError) Error() string {
	return fmt.Sprintf("error type %v - method %v - error message: %v", a.Type, a.Method, a.Message)
}

func (a *APIError) InvalidJSON(mt string, err error) *APIError {
	a.StatusCode = http.StatusBadRequest
	a.Type = "invalid-json"
	a.Method = mt
	a.Message = "Invalid or malformed JSON, " + err.Error()
	return a
}

func (a *APIError) BadRequest(mt string, err error) *APIError {
	a.StatusCode = http.StatusBadRequest
	a.Type = "client-error"
	a.Method = mt
	a.Message = "Cannot process current request, " + err.Error()
	return a
}

func (a *APIError) InternalServerError(mt string, err error) *APIError {
	a.StatusCode = http.StatusInternalServerError
	a.Type = "server-error"
	a.Method = mt
	a.Message = "Internal server error" + err.Error()
	return a
}
