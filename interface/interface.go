package main

import "fmt"

type shape interface {
	Area() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type triangle struct {
	length  float64
	breadth float64
	height  float64
}

func main() {
	c := circle{6.6}
	s := square{10.2}
	t := triangle{4.5, 5.3, 2.6}

	calculateArea(c)
	calculateArea(s)
	calculateArea(t)
}

func calculateArea(i shape) {
	fmt.Println(i.Area())
}

func (c circle) Area() float64 {
	return c.radius * c.radius
}

func (s square) Area() float64 {
	return s.length * s.length * s.length
}

func (t triangle) Area() float64 {
	return t.length * t.height * t.breadth
}
