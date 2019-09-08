package types

import (
	"errors"
	"net/http"
)

//	// 10xx Main App Error
//	AppServerError
//	InvalidLocale
//	InvalidTimezone
//	ExceedRequestLimit
//	// 11xx Http errors
//	InvalidHeaderOrParameter
//	HTTPNotFound
//	HTTPInternalError
//	// 12xx Auth Errors
//	AuthenticationFailure
//	AuthenticationParamMissing
//	UnauthorizedAccess
//	// 13xx Session Errors
//	SessionExpired
//	SessionInvalid
//	InvalidEmailOrPassword
var (
	// ErrNotFound -
	ErrNotFound            = errors.New("PAGE NOT FOUND")
	ErrInternalServerError = errors.New("INTERNAL SERVER ERROR")
)

type ErrorMessage struct {
	Code    string
	Message string
}

// ErrorCode is error type to used
type ErrorCode string

func (ec ErrorCode) String() string {
	return string(ec)
}

// Customed error codes
const (
	AppServerError     ErrorCode = "1001"
	InvalidLocale      ErrorCode = "1002"
	InvalidTimezone    ErrorCode = "1003"
	ExceedRequestLimit ErrorCode = "1004"

	HttpInvalidHeaderOrParam ErrorCode = "1101"
	HTTPNotFound             ErrorCode = "1102"
	HTTPInternalError        ErrorCode = "1103"

	AuthenticationFailure      ErrorCode = "1201"
	AuthenticationParamMissing ErrorCode = "1202"
	UnauthorizedAccess         ErrorCode = "1203"

	SessionExpired         ErrorCode = "1301"
	SessionInvalid         ErrorCode = "1302"
	InvalidEmailOrPassword ErrorCode = "1303"
)

var codeStatusMap = map[ErrorCode]int{
	AppServerError:     http.StatusInternalServerError,
	InvalidLocale:      http.StatusInternalServerError,
	InvalidTimezone:    http.StatusInternalServerError,
	ExceedRequestLimit: http.StatusInternalServerError,

	HttpInvalidHeaderOrParam: http.StatusBadRequest,
	HTTPNotFound:             http.StatusNotFound,
	HTTPInternalError:        http.StatusInternalServerError,

	AuthenticationFailure:      http.StatusForbidden,
	AuthenticationParamMissing: http.StatusForbidden,
	UnauthorizedAccess:         http.StatusForbidden,

	SessionExpired:         http.StatusForbidden,
	SessionInvalid:         http.StatusForbidden,
	InvalidEmailOrPassword: http.StatusForbidden,
}

// GetHTTPStatus returns http status mapped to customer error
func GetHTTPStatus(code ErrorCode) int {
	return codeStatusMap[code]
}

var codeMessageMap = map[ErrorCode]string{
	AppServerError:     "App server error ",
	InvalidLocale:      "Invalid locale",
	InvalidTimezone:    "Invalid timezone",
	ExceedRequestLimit: "Exceed max limit of reqeust",

	HttpInvalidHeaderOrParam: "Invalid headers or pamraters",
	HTTPNotFound:             "404, not found.",
	HTTPInternalError:        "Oops! Internal error, please try again.",

	AuthenticationFailure:      "Authorizated failed",
	AuthenticationParamMissing: "Authorizated failed",
	UnauthorizedAccess:         "Unauthorized Access",

	SessionExpired:         "Session expired",
	SessionInvalid:         "Session invalid",
	InvalidEmailOrPassword: "Invalid email or password",
}

// GetErrorMessage -
func GetErrorMessage(code ErrorCode) string {
	return codeMessageMap[code]
}
