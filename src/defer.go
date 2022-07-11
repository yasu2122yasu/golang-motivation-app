package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("World foo")

	fmt.Println("Hello Hoo")
}

func main() {
	file, _ := os.Open("./if.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
