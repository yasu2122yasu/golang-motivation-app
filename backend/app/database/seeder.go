package database

import (
	"app/model"
	"fmt"

	"gorm.io/gorm"
)

// 最終的には別の方法でSeedingを実装する。
func UserSeeder(db *gorm.DB) {
	// 詳細なエラーハンドリングの追加
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		fmt.Println("データベースは作成ずみ")
	}
	users := model.User{Id: 1, Name: "abe", Email: "mailmail@gmail.com", Password: "password"}

	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("%+v", err)
	}
}
