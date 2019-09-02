package fib

import "testing"

func TestFib0(t *testing.T) {
	var (
		in       = 0
		expected = 0
	)
	actual, _ := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}
func TestFib1(t *testing.T) {
	var (
		in       = 1
		expected = 1
	)
	actual, _ := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}
func TestFib2(t *testing.T) {
	var (
		in       = 2
		expected = 1
	)
	actual, _ := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}
func TestFib7(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual, _ := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}
func TestFibIllegarIInputError(t *testing.T) {
	var (
		in       = -1
		expected = "illegal integer input"
	)
	_, err := Fib(in)
	if err != nil && err.Error() != expected {
		t.Errorf("Fib(%d) get error '%s'; expected '%s'", in, err.Error(), expected)
	}
}
func TestFibOverflowError(t *testing.T) {
	var (
		in       = 50000
		expected = "integer overflow"
	)
	_, err := Fib(in)
	if err != nil && err.Error() != expected {
		t.Errorf("Fib(%d) get error '%s'; expected '%s'", in, err.Error(), expected)
	}
}
