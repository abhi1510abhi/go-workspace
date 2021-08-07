package main

import (
	"fmt"
)

func main() {

	var x int = 10
	var p *int = &x
	x = 15

	fmt.Println(x)
	fmt.Println(p)  // it will print the address;
	fmt.Println(*p) // it will print the value;

	// type notes

	type kelvin int
	type celcius int

	var a kelvin = 20
	var c celcius
	c = celcius(a) // type casting
	fmt.Println(a, c)

}
