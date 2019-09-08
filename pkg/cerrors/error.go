package cerrors

// GeneralError - public error
type GeneralError interface {
	Code() ErrorCode
	Messages() []string
}

// InternalError -
type InternalError interface {
	Internal() bool
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
