package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get Out!")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	DriveCar(mike)
	DriveCar(x)
}

//インターフェイスにあるメソッドを必ず使ってほしい時にインターフェイスを使う
/*
func main()の中の&Personにアドレスを渡す必要性がある

*/
