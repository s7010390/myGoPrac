package main

import "fmt"

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge
func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func (t triangularPrism) Volume() float64 {
	return 0.5 * t.base * t.attitude * t.height
}

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

func main() {
	var a interface{}
	a = 10
	fmt.Printf("%T %v\n", a, a)

	a = "ten"
	fmt.Printf("%T %v\n", a, a)

	a = true
	fmt.Printf("%T %v\n", a, a)

	a = func() string { return "hello" }
	fmt.Printf("%T %v\n", a, a)

	c := cube{edge: 3.8}
	t := triangularPrism{base: 4, attitude: 5, height: 6}
	cubeVolume := VolumeOf(c)
	triangularPrismVolume := VolumeOf(t)
	fmt.Printf("Volume of the cube: %.2f\n", cubeVolume)
	fmt.Printf("Volume of the triangular prism: %.2f\n", triangularPrismVolume)
}
