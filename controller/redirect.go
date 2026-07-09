package controller

import (
	"net/http"
	"time"
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

	if url.Original == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if time.Now().After(url.Expires) {
		c.JSON(http.StatusGone, gin.H{"error": "URL has expired"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Original)
}
