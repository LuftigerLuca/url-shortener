package main

import (
	"log"
	"url-shortener/controller"
	"url-shortener/persistence"

	"github.com/gin-gonic/gin"
)

func main() {
	persistence.ConnectDataBase()

	router := gin.Default()
	router.GET("/:short", controller.Redirect)

	urlCtrl := controller.UrlController{}

	api := router.Group("/api")
	{
		api.POST("/create", urlCtrl.CreateShortUrl)
		api.DELETE("/delete", urlCtrl.DeleteShortUrl)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("could not start router:", err)
	}
}
