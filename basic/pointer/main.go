package main

import "fmt"

func main() {
	d := 2
	double(&d)
	fmt.Print(d)
}
func double(d *int) {
	*d = *d * 2
}
