package types

// GeneralError - public error
type GeneralError interface {
	Code() ErrorCode
	Messages() []string
}

// InternalError -
type InternalError interface {
	Internal() bool
}

type dbError struct {
	// custom message with context information, i.e., [createUser]
	message string
	// store original error created by library
	originalError error
}

func (de *dbError) Internal() bool {
	return true
}
func (de *dbError) Error() string {
	return de.message + ":" + de.originalError.Error()
}
