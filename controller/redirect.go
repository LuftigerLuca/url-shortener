package controller

import (
	"net/http"
	"url-shortener/model"
	"url-shortener/persistence"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	short := c.Param("short")

	url := model.URL{}
	result := persistence.DB.Where("short = ?", short).Find(&url)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Original)
}
