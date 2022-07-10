package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	//Go言語の場合はもっと簡単に書ける、iはインデックス番号、vは値。使いたくない時は_を挿れる

	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//キーの時だけは_不要で、キーだけで取り出せる

	for t := range m {
		fmt.Println(t)
	}

	for _, v := range m {
		fmt.Println(v)
	}

}
