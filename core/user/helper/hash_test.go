package helper

import (
	"fmt"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	var (
		in = "admin"
	)
	actual := HashAndSalt([]byte("admin"))
	fmt.Println("actual", actual)
	if !ComparePassword(actual, []byte(in)) {
		//	if actual != expected {
		t.Errorf("HashAndSalt(%s) = %s;", in, actual)
	}
}
