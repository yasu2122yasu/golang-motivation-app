package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("peocess", i*1000)
		}()
	}
	fmt.Println("#######################")
}

/*
func main()内でsync.WaitGroup型のwgを作る
*/
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	//Addメソッドは終了するまで待ちたいゴルーチンの数だけAddを呼ぶ。
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
