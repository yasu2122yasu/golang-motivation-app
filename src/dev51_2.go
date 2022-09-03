//Buffered Channels

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	//forとrangeを回すときは、closeでchを一度閉じる必要がある。

	for c := range ch {
		fmt.Println(c)
	}
}
