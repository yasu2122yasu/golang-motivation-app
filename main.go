package main

import (
	"encoding/json"
	"fmt"
	"go_practice/database"
	"go_practice/repository"
	"go_practice/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// グローバル変数
var savingService service.IIndexSavingService

func NewRouting(data database.ConnectDatabase){
    // データベースのインスタンスを作成
	db, err := data.Connect()
	if err != nil {
		log.Fatalf("データベース接続失敗: %v", err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("データベース接続失敗: %v", err)
		return
	}
	fmt.Println("データベース接続成功")

	// リポジトリとサービスのインスタンスを作成
    savingService = service.NewIndexSavingService(repository.NewIndexSavingRepository(data))
}

func indexSaving(w http.ResponseWriter, r *http.Request) {
	iss, err := savingService.IndexSaving()

	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(iss)
}

func main() {
    data := database.NewConnectDatabase()

    NewRouting(data)

	// ルーティング
	r := mux.NewRouter()
	r.HandleFunc("/", indexSaving).Methods("GET")
	// 他のルート定義もこちら...

	log.Fatal(http.ListenAndServe(":8000", r))
}
