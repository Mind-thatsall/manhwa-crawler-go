package main

import "github.com/Mind-thatsall/web-crawler-go/models"

func main() {
	// r := gin.Default()

	// r.GET("/api/manhwas", controller.FindManhwas)
	// r.GET("/api/manhwa/:id", controller.FindManhwa)
	// r.GET("/api/manhwa/series/:id")

	// r.Run("localhost:9090")

	models.GetChapter("starting-from-today-ill-work-as-a-city-lord-chapter-108")
}
