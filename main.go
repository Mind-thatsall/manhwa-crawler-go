package main

import (
	"github.com/Mind-thatsall/web-crawler-go/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/manhwas", controller.FindManhwas)
	r.GET("/manhwas/:id", controller.FindManhwa)

	r.Run("localhost:9090")

	//models.GetAllChapters("the-first-sword-of-earth")
}
