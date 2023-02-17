package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	manhwas := getInfos()

	db.AutoMigrate(&Manhwa{})
	db.Create(&manhwas)

	DB = db
}
