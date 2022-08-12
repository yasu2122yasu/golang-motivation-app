package main 

import(
	"fmt"
)

type UserNotFound struct {
	Username string
}

func myFunc() error {
	return nil
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}