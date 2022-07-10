package main

import "fmt"

//可変長変数なので、引数の数が変わっても問題ない！
func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}

}

func main() {
	foo()
	foo(10, 20)
	foo(10, 20, 30)
	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...)
}
