package main

import (
	"log/slog"
	"net/http"
	"url-shortener/controller"
	"url-shortener/persistence"
	"url-shortener/service"
)

func main() {
	DB := persistence.ConnectDataBase()
	mux := http.NewServeMux()

	urlService := service.UrlService{
		DB: DB,
	}

	redirectCtrl := controller.RedirectController{
		DB: DB,
	}

	urlCtrl := controller.UrlController{
		Service: urlService,
		DB:      DB,
	}

	mux.HandleFunc("/{short}", redirectCtrl.Redirect)
	mux.HandleFunc("/api/create", urlCtrl.CreateShortUrl)
	mux.HandleFunc("/api/delete", urlCtrl.DeleteShortUrl)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Could not start the webserver:", err.Error())
	}
}
