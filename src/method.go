//39の場合、アスタリスクを使うことで実体を書き換えることができる

package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

//アスタリスクを利用することでVertexの中身を変えることができる。
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Area())
}
