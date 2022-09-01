//チャネルの書き方
package main

import (
	"fmt"
)

/*
_は（アンダースコア変数）、宣言はするものの後でも使わない変数のこと。
この場合、配列のインデックスである0,1,2,3などを出力しないために使用
されている。

チャネルは送る方向に矢印を指す。c <- sumのように。
*/
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

/*
make()、make関数とは組み込み関数の1つで、makeを使用することでsliceを指定することができる。
例：make([]T, len, cap)
[]Tはスライスの要素の型を指定し、
!スライスの長さがスライスの容量を超えることはない。

以下のように利用することもある
チャネル変数 := make(chan 送受信するデータ型)
*/

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine1(s, c)
	x := <-c
	fmt.Println(x)
}
