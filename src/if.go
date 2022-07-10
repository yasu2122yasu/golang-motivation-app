package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	/*簡略化したのが22行目から
	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}
	*/

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("Great-2")
	}

	num := 4
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("else")
	} else {
		fmt.Println("DDD")
	}

	x, y := 10, 20
	if x == 10 && y == 20 {
		fmt.Println("&&")
	}

	if x == 10 || y == 20 {
		fmt.Println("SFF")
	}

}
