package model

import "time"

type User struct {
	// todo: idが自動増分されるやつにしたい
	Id       int    `gorm:"id"`
	Name     string `gorm:"name"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}

// User は複数の言語を所有し、かつ言語に属しています。`user_languages` が結合テーブルになります
type Country struct {
	Id        int        `gorm:"id"`
	Languages []Language `gorm:"many2many:country_languages;"`
}

type Language struct {
	Id   int    `gorm:"id"`
	Name string `gorm:"name"`
}

type DateMaster struct {
	Date      time.Time `gorm:"primaryKey"` // 日付を主キーとして使用
	DayOfWeek int       // 曜日
	IsWeekend bool      // 週末かどうか
	IsHoliday bool      // 祝日かどうか
	Quarter   int       // 四半期
	Year      int       // 年
	Month     int       // 月
}
