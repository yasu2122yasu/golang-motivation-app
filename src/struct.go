package main

import "fmt"

//構造体は大文字を利用する。
type Vertex struct {
	// X Y int　　このように書くこともできる
	X int
	Y int
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	//指定がなければYにはデフォルトのintegerの0が入る
	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	//v4では空の出力をする
	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	//ここでは別の書き方をしてみる、Vertexという構造体の代入をまとめてやっている
	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	//v6は明示的にポインタが帰ってくるとは分かりづらい、v7と同じようにポイントを返している。スライス、マップではv6を利用するケースが多い
	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

}
