package controller

import (
	"net/http"

	"github.com/Mind-thatsall/web-crawler-go/models"
	"github.com/gin-gonic/gin"
)

func FindManhwas(c *gin.Context) {
	var manhwas []models.Manhwa
	models.DB.Find(&manhwas)

	c.JSON(http.StatusOK, gin.H{"data": manhwas})
}
