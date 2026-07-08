package controller

import (
	"net/http"
	"url-shortener/config"
	"url-shortener/model"
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	Service service.UrlService
}

func (ctrl *UrlController) CreateShortUrl(c *gin.Context) {
	var req model.CreateUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, created, err := ctrl.Service.CreateShortUrl(req)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	res := model.CreateUrlResponse{
		Short:    url.Short,
		ShortUrl: config.BaseUrl + "/" + url.Short,
	}

	if created {
		c.JSON(http.StatusCreated, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (ctrl *UrlController) DeleteShortUrl(c *gin.Context) {
	var req model.DeleteUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Service.DeleteShortUrl(req)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted URL"})
}
