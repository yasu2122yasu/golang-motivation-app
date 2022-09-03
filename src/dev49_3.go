package main

import (
	"fmt"
	"sync"
	"time"
)

/*この場合だと func main()で実行されたqg.Wait()をfunc goroutine()内
でwg.Doneで終わらせているからfunc normal() が出力されない。
wg.Doneを書かないとエラーが発生する。
*/
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
