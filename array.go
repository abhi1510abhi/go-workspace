package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array Demo printing 1 to 100")
	var num [100]int

	for i := 0; i < 100; i++ {
		num[i] = i
	}
	for j := 0; j < 100; j++ {
		fmt.Println(num[j])
	}

	fmt.Println("length :", len(num))
}
