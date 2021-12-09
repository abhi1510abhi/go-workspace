package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Circle struct {
	radius float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) myName() {
	fmt.Println("Rectangle")
}
func main() {

	c := Circle{3}
	r := Rectangle{3, 4}
	//fmt.Println(c.area())
	//fmt.Println(r.area())

	// slice of shapes include circle and rectangle struct
	s := []Shape{c, r}

	for _, ele := range s {
		// looping over and printing area of all shapes
		fmt.Println(ele.area())
	}
	r.myName()
}
