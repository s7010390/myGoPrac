package main

import "fmt"

type rectangle struct {
	w, l float64
}

func (rectang rectangle) Area() float64 {
	return rectang.l * rectang.w
}
func (rectang *rectangle) setW(w float64) {
	rectang.w = w
}
func (rectang *rectangle) setL(l float64) {
	rectang.l = l
}

func main() {
	rect := rectangle{
		w: 4,
		l: 5}
	rect.l = 20
	rect.setW(52)
	fmt.Printf("w: %.2f, l:%.2f, area:%f", rect.w, rect.l, rect.Area())

}
