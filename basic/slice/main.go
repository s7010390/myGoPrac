package main

import "fmt"

func main() {
	var s []int
	if s == nil {
		fmt.Println("is nil !!")
	}

	a := [...]int{1, 2, 3, 4}
	s = a[:]
	fmt.Printf("%d, %d, %p, %p\n", len(s), cap(s), s, &a)
	s = append(s, 1, 2)
	fmt.Printf("%d, %d, %p, %p\n", len(s), cap(s), s, &a)

}
