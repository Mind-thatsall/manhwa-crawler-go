package main

import (
	"github.com/Mind-thatsall/web-crawler-go/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin,access-control-allow-headers"},
	}))

	dataRoutes := r.Group("api/manhwas")
	{
		dataRoutes.GET("/:page", controller.FindManhwas)
		dataRoutes.GET("/series/:id", controller.FindManhwa)
		dataRoutes.GET("/series/chapters/:id", controller.FindChapter)
	}

	r.Run("localhost:9090")

	// models.GetManhwas()
}
