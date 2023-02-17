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

func FindManhwa(c *gin.Context) {
	var manhwa models.ManhwaData
	if err := models.DB.Where("slug = ?", c.Param("id")).First(&manhwa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": manhwa})
}
