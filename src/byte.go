package main

import "fmt"

func main() {
	b := []byte{71, 72}
	fmt.Println(b)

	//アスキーコードとして表示
	fmt.Println(string(b))

	c := []byte("Hi!")
	fmt.Println(c)
	fmt.Println(string(c))
}
