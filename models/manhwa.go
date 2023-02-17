package models

import "gorm.io/gorm"

type Manhwa struct {
	ID      string `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	Slug    string `json:"slug"`
}

type ManhwaData struct {
	gorm.Model
	ID          string    `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Chapters    []Chapter `json:"chapters" gorm:"foreignKey:Slug"`
	Slug        string    `json:"slug"`
}

type ManhwaChapter struct {
}

type Chapter struct {
	ID     string `json:"id" gorm:"primary_key"`
	Number string `json:"number"`
	Date   string `json:"date"`
	Slug   string `json:"slug"`
}
