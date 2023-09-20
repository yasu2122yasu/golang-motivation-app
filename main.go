package main

import (
	"fmt"
	"go_practice/database"
	"go_practice/repository"
)

func main() {
    db := database.Connect()
    defer db.Close()

    err := db.Ping()

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }

		repo, err := repository.GetPersons()
		// if err != nil{
		// 	log.Fatal()
		// }
		fmt.Println(repo)
}
