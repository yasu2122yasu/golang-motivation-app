package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type IConnectDatabase interface {
    Connect() (*sql.DB, error)
}

type ConnectDatabase struct{}

func NewConnectDatabase() ConnectDatabase {
    return ConnectDatabase{}
}

func (c ConnectDatabase) Connect() (*sql.DB, error) {
    err := godotenv.Load()
    if err != nil {
        fmt.Println(err.Error())
    }


    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    // host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    database_name := os.Getenv("DB_DATABASE_NAME")

    dbconf := user + ":" + password + "@tcp(" + ":" + port + ")/" + database_name + "?charset=utf8mb4"
    db, err := sql.Open("mysql", dbconf)
    if err != nil {
        fmt.Println(err.Error())
    }
    return db,err
}
