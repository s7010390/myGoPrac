package main

import (
	"fmt"
	"time"
)

func doSometing(i int, tra int) int {
	if tra == 0 {
		fmt.Printf("result: %d\n", i)
		return i
	}
	time.Sleep(time.Second)
	sum := i + i
	tra -= 1
	return doSometing(sum, tra)
}
func main() {
	fmt.Printf("start\n")
	go doSometing(1, 3)
	go doSometing(1, 10)
	go doSometing(1, 1)

	time.Sleep(15 * time.Second)
	fmt.Printf("end\n")
}

//lightweigth thread
