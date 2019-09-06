package types

import (
	"errors"
)

var (
	// ErrPageNotFound -
	ErrNotFound            = errors.New("PAGE NOT FOUND")
	ErrInternalServerError = errors.New("INTERNAL SERVER ERROR")
)

type ErrorMessage struct {
	Code    string
	Message string
}

//var errorMessages []*ErrorMessage
//var inited = false
//
//func GetErrorMessages() []*ErrorMessage {
//	if inited == false {
//		// 10xx Main App Error
//		append(errorMessages, &ErrorMessage{"1000", "App Server Error, please contect the admin."})
//		append(errorMessages, &ErrorMessage{"1001", "Missing Headers"})
//		append(errorMessages, &ErrorMessage{"1002", "Missing Parameters"})
//		append(errorMessages, &ErrorMessage{"1003", "Invalid offset or limit"})
//		append(errorMessages, &ErrorMessage{"1004", "Invalid Locale"})
//		append(errorMessages, &ErrorMessage{"1005", "Invalid Timezone"})
//		append(errorMessages, &ErrorMessage{"1006", "Exceed the limit of requests"})
//
//		// 11xx Http errors
//		append(errorMessages, &ErrorMessage{"1101", "Unauthorized"})
//		append(errorMessages, &ErrorMessage{"1102", "Not authorized to access"})
//		append(errorMessages, &ErrorMessage{"1103", "Authentication Failed"})
//		append(errorMessages, &ErrorMessage{"1104", "Not Found"})
//
//		// 12xx Auth Errors
//		append(errorMessages, &ErrorMessage{"1201", "Your session is expired, please login again"})
//		append(errorMessages, &ErrorMessage{"1202", "Your session is invalid"})
//		append(errorMessages, &ErrorMessage{"1203", "Your session token is invalid"})
//		append(errorMessages, &ErrorMessage{"1204", "You are Unauthorized, please login"})
//		append(errorMessages, &ErrorMessage{"1205", "Authentication Error, User Not Found"})
//
//		// 13xx Session Errors
//		append(errorMessages, &ErrorMessage{"1301", "Invalid Login Type"})
//		append(errorMessages, &ErrorMessage{"1302", "Login Error"})
//		append(errorMessages, &ErrorMessage{"1303", "Your Account is disabled by the admin"})
//		append(errorMessages, &ErrorMessage{"1304", "Invalid Mobile Numbers"})
//		append(errorMessages, &ErrorMessage{"1305", "Invalid Email or Password"})
//		append(errorMessages, &ErrorMessage{"1306", "Wrong confirmation code! Try Again"})
//
//		init = true
//	}
//	return errorMessages
//}
