package main

import (
	"app/controller"
	"app/database"
	"app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	database.ConnectDatabase()
	// Ginのルーターの初期化
	r := gin.Default()
	// フロントエンドとの疎通に必要なCORSの設定
	r.Use(middleware.Cors())

	r.GET("/users", controller.GetAllUser)
	r.GET("/users/:id", controller.GetOneUser)
	r.POST("/users", controller.CreateUser)
	r.Run(":8080")
}
