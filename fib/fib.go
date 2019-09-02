package fib

import "errors"

const (
	// MinUint binary all zero
	MinUint uint = 0
	// MaxUnit binary all one
	MaxUnit = ^MinUint
	// MaxInt >> 1
	MaxInt = int(MaxUnit >> 1)
	// MinInt ^
	MinInt = ^MaxInt
)

var (
	// ErrIllegalInput for improper input
	ErrIllegalInput = errors.New("illegal integer input")
	// ErrOverFlow for integer overflow
	ErrOverFlow = errors.New("integer overflow")
)

// Fib generate fib number to given input number
func Fib(n int) (int, error) {
	if n < 2 && n >= 0 {
		return n, nil
	} else if n >= 2 {
		a, b := 0, 1
		for i := 0; i < n-1; i++ {
			a, b = b, a+b
		}
		if b < 0 {
			return -1, ErrOverFlow
		}
		return b, nil
	}
	return -1, ErrIllegalInput
}
