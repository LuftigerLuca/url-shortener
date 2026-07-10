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

	if err := http.ListenAndServe(":8080", middleware(mux)); err != nil {
		slog.Error("Could not start the webserver:", "error", err.Error())
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				slog.Warn("web middleware recovery caught smth", "error", err)
				http.Error(w, "There was an internal server error", http.StatusInternalServerError)
			}
		}()

		slog.Info(r.URL.Path + " called from " + r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
