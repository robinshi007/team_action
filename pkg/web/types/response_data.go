package types

type ResponseData struct {
	ErrorCode int         `json:"error_code,omitempty"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
