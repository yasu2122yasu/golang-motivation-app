package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	v, ok := m["nothing"]
	fmt.Println(v, ok)

	v, ok = m["apple"]
	fmt.Println(v, ok)
	//v＝value、100 trueが出力される

	v = m["banana"]
	fmt.Println(v)

	//okがない場合、true,falseの戻り値がなくなる

	//キーがstring型、値がint型、map上に空のメモリを作ってから代入することも可能
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}

}
