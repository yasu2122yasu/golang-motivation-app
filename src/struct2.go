package main

import "fmt"

//構造体は大文字を利用する。36
type Vertex struct {
	// X Y int　　このように書くこともできる
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

//ポインタを使わないと書き換えができないイメージ？
func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(*v2)
}
