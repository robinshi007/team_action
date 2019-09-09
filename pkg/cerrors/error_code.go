package cerrors

import (
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

// Customed error codes
const (
	AppServerError     string = "1001"
	InvalidLocale      string = "1002"
	InvalidTimezone    string = "1003"
	ExceedRequestLimit string = "1004"

	HTTPInvalidHeaderOrParam string = "1101"
	HTTPNotFound             string = "1102"
	HTTPInternalError        string = "1103"

	AuthenticationFailure      string = "1201"
	AuthenticationParamMissing string = "1202"
	UnauthorizedAccess         string = "1203"

	SessionExpired         string = "1301"
	SessionInvalid         string = "1302"
	InvalidEmailOrPassword string = "1303"
)

var codeStatusMap = map[string]int{
	AppServerError:     http.StatusInternalServerError,
	InvalidLocale:      http.StatusInternalServerError,
	InvalidTimezone:    http.StatusInternalServerError,
	ExceedRequestLimit: http.StatusInternalServerError,

	HTTPInvalidHeaderOrParam: http.StatusBadRequest,
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
func GetHTTPStatus(code string) int {
	return codeStatusMap[code]
}

var codeMessageMap = map[string]string{
	AppServerError:     "App server error ",
	InvalidLocale:      "Invalid locale",
	InvalidTimezone:    "Invalid timezone",
	ExceedRequestLimit: "Exceed max limit of reqeust",

	HTTPInvalidHeaderOrParam: "Invalid headers or pamraters",
	HTTPNotFound:             "404, not found.",
	HTTPInternalError:        "Oops! Something went wrong, please try again.",

	AuthenticationFailure:      "Authorizated failed",
	AuthenticationParamMissing: "Authorizated failed",
	UnauthorizedAccess:         "Unauthorized Access",

	SessionExpired:         "Session expired",
	SessionInvalid:         "Session invalid",
	InvalidEmailOrPassword: "Invalid email or password",
}

// GetErrorMessage -
func GetErrorMessage(code string) string {
	return codeMessageMap[code]
}
