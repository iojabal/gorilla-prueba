package models

import (
	"prueba/db"
)

type Article struct {
	//gorm.Model
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	UserId      int64  `gorm:"foreignKey:user_id"`
	User        User   // `gorm:"column:id"`
}

type Articles []Article

func MigrateArt() {
	db.Database.AutoMigrate(Article{})
}
