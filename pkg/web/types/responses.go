package types

// SuccessResponse -
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse -
type ErrorResponse struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}
