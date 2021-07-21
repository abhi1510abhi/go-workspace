package main

import (
	"fmt"
)

func calculate(num1 float64, num2 float64, operator string) float64 {
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	}

	return result
}

func main() {
	fmt.Println("Calculator")

	var num1, num2, result float64
	var operator string
	fmt.Println("Number 1")
	fmt.Scanln(&num1)

	fmt.Println("Number 2")
	fmt.Scanln(&num2)

	fmt.Println("operator")
	fmt.Scanln(&operator)

	result = calculate(num1, num2, operator)
	fmt.Println(result)

}
