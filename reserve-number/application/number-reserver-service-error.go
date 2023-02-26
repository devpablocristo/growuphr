package application

// import (
// 	"fmt"
// )

// type NumberReserverError struct {
// 	Type       string `json:"type,omitempty"`
// 	Method     string `json:"method,omitempty"`
// 	MetPak     string `json:"package,omitempty"`
// 	Message    string `json:"message,omitempty"`
// }

// func NewAPIError(st int, tp string, mt string, pk string, mg error) *APIError {
// 	return &APIError{
// 		StatusCode: st,
// 		Type:       tp,
// 		Method:     mt,
// 		MetPak:     pk,
// 		Message:    mg.Error(),
// 	}
// }

// func (a *APIError) Error() string {
// 	return fmt.Sprintf("status code: %v, error type: %v - method: %v.%v - error message: %v", a.StatusCode, a.Type, a.MetPak, a.Method, a.Message)
// }
