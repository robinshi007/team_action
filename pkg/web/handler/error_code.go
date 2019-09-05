package handler

import (
	"errors"
)

var (
	// ErrPageNotFound -
	ErrNotFound            = errors.New("PAGE NOT FOUND")
	ErrInternalServerError = errors.New("INTERNAL SERVER ERROR")
)
