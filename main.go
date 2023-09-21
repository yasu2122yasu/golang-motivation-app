package main

import (
	"encoding/json"
	"fmt"
	"go_practice/database"
	"go_practice/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexSaving(w http.ResponseWriter, r *http.Request) {
    iss, err := service.IndexSavingService()

    if err != nil {
        log.Fatal()
    }

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(iss)

}

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
    r := mux.NewRouter()

    r.HandleFunc("/", indexSaving).Methods("GET")
    // r.HandleFunc("/api/savings/{id}", ).Methods("GET")
    // r.HandleFunc("/api/books", createBook).Methods("POST")
    // r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
    // r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", r))



}
