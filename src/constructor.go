package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.y * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func Area(v Vertex) int {
	return v.x * v.y
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Area())
}

//パッケージ名.Newでstructを返す。例：log.New
