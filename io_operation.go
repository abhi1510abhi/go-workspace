package main

import (
	"fmt"
	"os"
)

func createFile() {

	// create file
	file, err := os.Create("myfile")
	if err != nil {
		fmt.Println(err)
	}
	// data can be passed only in byte []
	data := []byte("hello world\n")
	file.Write(data)
	//close the file once operation is done
	file.Close()
}

func readFile() {
	//open the file
	file, err := os.Open("myfile")
	if err != nil {
		fmt.Println(err)
	}
	// create byte [] to receive data
	data := make([]byte, 100)

	// read file
	file.Read(data)

	fmt.Println("Reading file")
	//print data
	fmt.Println(string(data))

}

func main() {

	createFile()
	fmt.Println("file created")
	readFile()
}
