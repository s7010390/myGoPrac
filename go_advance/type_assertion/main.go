package main

import "fmt"

func main() {
	var a, b Obj
	a = rectangle{l: 5, w: 5}
	b = triangle{base: 4, h: 4}

	//convention

	fmt.Printf("a.area(): %f\n", a.Area())
	fmt.Printf("b.area(): %f\n", b.Area())
	if v, ok := b.(triangle); ok {
		v.h = 20
		fmt.Printf("v.area(): %f\n", v.Area())
	}

}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

type triangle struct {
	base, h float64
}

func (rec rectangle) Area() float64 {
	return rec.w * rec.l
}

func (tri triangle) Area() float64 {
	return tri.base * tri.h * 0.5
}
