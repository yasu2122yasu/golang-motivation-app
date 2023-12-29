package model

type User struct {
	// todo: idが自動増分されるやつにしたい
	Id       int    `gorm:"id"`
	Name     string `gorm:"name"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}
