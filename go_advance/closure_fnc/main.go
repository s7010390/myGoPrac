package main

import "fmt"

func main() {
	fn_1 := newCounterFunc()
	fn_2 := newCounterFunc()
	fmt.Printf("fn: %d\n", fn_1())
	fmt.Printf("fn: %d\n", fn_1())
	fmt.Printf("fn: %d\n", fn_2())
	fmt.Printf("fn: %d\n", fn_1())
	fmt.Printf("fn: %d\n", fn_2())
}

func newCounterFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
