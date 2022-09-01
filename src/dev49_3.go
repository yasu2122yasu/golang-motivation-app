package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("World", &wg)
	normal("hello")
	wg.Wait()
}
