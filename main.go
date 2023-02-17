package main

import (
	"github.com/Mind-thatsall/web-crawler-go/controller"
	"github.com/Mind-thatsall/web-crawler-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectToDB()

	router.GET("/api/manhwas", controller.FindManhwas)

	router.Run("localhost:9090")

}
