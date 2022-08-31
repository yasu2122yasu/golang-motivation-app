package main

import (
	"fmt"
)

type Ver struct {
	X, Y int
}

func (v Ver) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
}

func main() {
	v := Ver{3, 4}
	fmt.Println(v)
}
