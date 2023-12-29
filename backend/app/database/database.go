package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	dsn := "host=database port=5432 user=postgres password=postgres dbname=app sslmode=disable"
	dialector := postgres.Open(dsn)
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		fmt.Println("db error!!")
		panic(err)
	}
	fmt.Println("db connected!!")
	Migration(Db)
	UserSeeder(Db)
}
