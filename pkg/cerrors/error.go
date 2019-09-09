package cerrors

// GeneralError - public error
type GeneralError interface {
	Code() string
	Messages() []string
	Error() string
}

// InternalError -
type InternalError interface {
	Internal() bool
}

// CustomError struct
type CustomError struct {
	ErrorCode string   `json:"code"`
	Errors    []string `json:"errors"`
}

// NewCustomError -
func NewCustomError(code string, errs []string) *CustomError {
	return &CustomError{
		ErrorCode: code,
		Errors:    errs,
	}
}

// Code -
func (pe *CustomError) Code() string {
	return pe.ErrorCode
}

// Messages -
func (pe *CustomError) Messages() []string {
	return pe.Errors
}

// Error -
func (pe *CustomError) Error() string {
	return pe.ErrorCode
}

// ParamError -
type ParamError struct {
	ErrorCode string   `json:"code"`
	Errors    []string `json:"errors"`
}

// NewParamError -
func NewParamError(errs []string) *ParamError {
	return &ParamError{
		ErrorCode: "1101",
		Errors:    errs,
	}
}

// Code -
func (pe *ParamError) Code() string {
	return pe.ErrorCode
}

// Messages -
func (pe *ParamError) Messages() []string {
	return pe.Errors
}

// Error --
func (pe *ParamError) Error() string {
	return pe.ErrorCode
}

// DBError -
type DBError struct {
	// custom message with context information, i.e., [createUser]
	message string
	// store original error created by library
	originalError error
}

// Internal -
func (de *DBError) Internal() bool {
	return true
}

// Error -
func (de *DBError) Error() string {
	return de.message + ":" + de.originalError.Error()
}
