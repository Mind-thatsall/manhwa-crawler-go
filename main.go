package main

import (
	"github.com/Mind-thatsall/web-crawler-go/controller"
	"github.com/Mind-thatsall/web-crawler-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	/* routes */
	r.GET("/manhwas", controller.FindManhwas)
	r.GET("/manhwas/:id", controller.FindManhwa)

	r.Run("localhost:9090")
}
