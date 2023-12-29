package database

import (
	"app/model"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	Db.AutoMigrate(&model.User{})
}
