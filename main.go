package main

import (
	"fmt"
	"reflect"
	"team_action/fib"
)

func main() {
	fmt.Println(reflect.TypeOf(1).String())
	printFib(-1)
	printFib(50)
	printFib(50000)
}

func printFib(input int) {
	result, err := fib.Fib(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Fib(%d) = %d\n", input, result)
}
