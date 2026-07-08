package controller

import (
	"net/http"
	"url-shortener/config"
	"url-shortener/model"
	"url-shortener/persistence"
	"url-shortener/util"

	"github.com/gin-gonic/gin"
)

type CreateUrlRequest struct {
	Url      string `json:"url" binding:"required"`
	Lifespan *int   `json:"lifespan"`
}

type CreateUrlResponse struct {
	Short    string `json:"short"`
	ShortUrl string `json:"short_url"`
}

func CreateShortUrl(c *gin.Context) {
	var req CreateUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := model.URL{}

	result := persistence.DB.Where("original = ?", req.Url).Find(&url)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected > 0 {
		res := CreateUrlResponse{}
		res.Short = url.Short
		res.ShortUrl = config.BaseUrl + "/" + url.Short

		c.JSON(http.StatusCreated, res)
		return
	}

	url.Original = req.Url
	url.Short = util.GenerateShortURL(8)

	if req.Lifespan != nil {
		url.Lifespan = *req.Lifespan
	} else {
		url.Lifespan = config.DefaultLifespan
	}

	if err := persistence.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := CreateUrlResponse{}
	res.Short = url.Short
	res.ShortUrl = config.BaseUrl + "/" + url.Short

	c.JSON(http.StatusCreated, res)
}

type DeleteUrlRequest struct {
	Short string `json:"short" binding:"required"`
}

func DeleteShortUrl(c *gin.Context) {
	var req DeleteUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := model.URL{}
	result := persistence.DB.Where("short = ?", req.Short).Delete(&url)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted URL"})
}
