package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getManhwas(c *gin.Context) {
	manhwas := getInfos()
	c.IndentedJSON(http.StatusOK, manhwas)
}

func main() {
	router := gin.Default()
	router.GET("/api/manhwas", getManhwas)

	router.Run("localhost:9090")

}
