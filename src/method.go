//39の場合、アスタリスクを使うことで実体を書き換えることができる。今回の内容はレシーバー

package main

import "fmt"

type Vertex struct {
	X, Y int
}

//値レシーバーの書き方。関数名の前がレシーバーで、関数の後が引数、戻り値の型となる
func (v Vertex) Area() int {
	return v.X * v.Y
}

//アスタリスクを利用することでVertexの中身を変えることができる。！ポイントレシーバー
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
