package models

type Manhwa struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	Slug    string `json:"slug"`
}

type ManhwaData struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	Chapters    []Chapter `json:"chapters"`
}

type ManhwaChapter struct {
}

type Chapter struct {
	ID     string `json:"id"`
	Number string `json:"number"`
	Date   string `json:"date"`
	Slug   string `json:"slug"`
}
