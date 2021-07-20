package main

import (
	"fmt"
	"log"
	"math"
)

var globalVariable int = 10 // global variable

func Add(x, y int) int { // Add -> can be access from outside package and add cannot be access from outside package
	return x + y
}

func main() {

	fmt.Println("Hello, world1")
	/**  Commment */
	// comment 2

	// getting user input
	var name string
	fmt.Println("Enter name")
	//fmt.Scanln(&name)
	fmt.Println(name)

	// variable decleration
	var x int = 12 // local variable
	var y int = 20
	var a, b int
	length := 1 // short hand of decleraing variable
	fmt.Println(x, y, a, b)

	// codition

	if (x == 11 && y == 20) || length == 1 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	//Loop

	for length = 1; length <= 5; length++ {
		fmt.Println("yo", length)
	}

	// functions
	result := Add(2, 3)
	fmt.Println("result", result)

	// using imported function
	var num float64 = 12
	var r = math.Sqrt(num)
	fmt.Println(math.Floor(r*1000) / 1000) // to roundoff till 2 decimal

	log.Print("This is a logger")

}
