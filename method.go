package main

import "fmt"

type Vertex struct {
	X float64
	Y float64
}

// add addition and subtraction with poins

func (v Vertex) addition(d float64) Vertex {
	v.X = v.X + d
	v.Y = v.Y + d
	return v
}

// it will change data since it is call by reference

func (v *Vertex) sub(d float64) {
	v.X = v.X - d
	v.Y = v.Y - d
}

func main() {

	v := Vertex{3, 4}
	fmt.Println(v)
	v = v.addition(5)
	fmt.Println(v)
	v.sub(5)
	fmt.Println(v)
}
