package main

import "fmt"

func do(i interface{}) {
	ii := i.(int)
	ii *= 2
	fmt.Println(ii)
}

func main() {
	do(10)
}

/*
interfaceはどんな型も許容できるが関数に代入する時には型変換を行う必要がある
interfaceを代入して計算するとエラーが出る
*/
