package main

import (
	"github.com/Mind-thatsall/web-crawler-go/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/manhwas/:page", controller.FindManhwas)
	r.GET("/api/manhwa/:id", controller.FindManhwa)
	r.GET("/api/manhwa/series/:id", controller.FindChapter)

	r.Run("localhost:9090")

	// models.GetManhwas()
}
