package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectToDB() {
	database, err := gorm.Open("sqlite3", "sqlite.db")

	if err != nil {
		log.Fatal(err)
	}

	manwhas := getInfos()

	database.AutoMigrate(&Manhwa{})
	database.Create(&manwhas)

	DB = database
}
