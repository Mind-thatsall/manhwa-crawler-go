package models

type Manhwa struct {
	ID      string `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	Slug    string `json:"slug"`
}
