//Buffered Channels

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	x := <-ch
	fmt.Println(x)

	y := <-ch
	fmt.Println(y)

	fmt.Println(len(ch))

	ch <- 300
	fmt.Println(len(ch))
}

/*
チャネルのバッファ数は次のように記述する。
c := make(chan 型名, バッファ数)


*/
