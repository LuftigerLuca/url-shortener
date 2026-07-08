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

	api := router.Group("/api")
	{
		api.POST("/create", controller.CreateShortUrl)
		api.DELETE("/delete", controller.DeleteShortUrl)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("could not start router:", err)
	}
}
