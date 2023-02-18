package controller

import (
	"net/http"

	"github.com/Mind-thatsall/web-crawler-go/models"
	"github.com/gin-gonic/gin"
)

func FindManhwas(c *gin.Context) {
	var manhwas []models.Manhwa
	manhwas = models.GetManhwas()

	c.JSON(http.StatusOK, gin.H{"data": manhwas})
}

func FindManhwa(c *gin.Context) {
	var manhwa models.ManhwaData
	manhwa = models.GetManhwaData(c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": manhwa})
}
