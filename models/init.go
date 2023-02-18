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

	db.Debug().AutoMigrate(&Manhwa{}, &ManhwaData{}, &Chapter{})

	manhwas := GetManhwas()

	db.Create(&manhwas)

	for i := 0; i < len(manhwas); i++ {
		manhwa := GetManhwaData(manhwas[i].Slug)
		db.Create(&manhwa)

		chaptersByManhwa := GetAllChapters(manhwas[i].Slug)
		for i := 0; i < len(chaptersByManhwa); i++ {
			db.Create(&chaptersByManhwa[i])
		}
	}

	DB = db
}
