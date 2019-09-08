package types

import (
	"team_action/pkg/cerrors"
)

// ResponseData -
type ResponseData struct {
	ErrorCode int         `json:"error_code,omitempty"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// SuccessResponse -
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse -
type ErrorResponse struct {
	Code   cerrors.ErrorCode `json:"code"`
	Errors []string          `json:"errors"`
}
