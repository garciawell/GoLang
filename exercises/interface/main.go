package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangule struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	tri := triangule{
		base:   12,
		height: 14,
	}
	sq := square{
		sideLength: 10,
	}

	printArea(tri)
	printArea(sq)

}

func (t triangule) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(val shape) {
	fmt.Println(val.getArea())
}
