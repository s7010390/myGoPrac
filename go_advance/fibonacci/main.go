package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(fibo(i), " ")
	}
}

func fibo(tra int) int {
	if tra <= 0 {
		return 0
	} else if tra == 1 {
		return 1
	} else {
		return fibo(tra-1) + fibo(tra-2)
	}
}
