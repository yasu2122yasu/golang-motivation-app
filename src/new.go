//makeとnewの違いはポインタを返すかどうか？

package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p *int = new(int)
	fmt.Printf("%T\n", p)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

}
