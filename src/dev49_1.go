package main

import (
	"fmt"
	"time"
)

/*
時間差でHelloを出力し続ける。
*/
func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	normal("hello")
}
