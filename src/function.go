package main

import "fmt"

func a() {
	fmt.Println("add function")
}

func b(x int, y int) {
	fmt.Println("addd function")
	fmt.Println(x + y)
}

func c(x int, y int) int {
	return x + y
}

//戻り値が2つある場合
func e(x, y int) (int, int) {
	return x + y, x - y
}

//引数の型が同じ場合、省略できる。帰り値で値を宣言すれば関数内で:=を利用しなくても良い

func cal(price, item int) (result int) {
	result = price * item
	//この場合、下記のresuletは不要
	return result
}

func main() {
	a()
	b(10, 20)
	d := c(10, 20)
	fmt.Println(d)

	e1, e2 := e(10, 20)
	fmt.Println(e1, e2)

	e3 := cal(100, 2)
	fmt.Println(e3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}

	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)

}
