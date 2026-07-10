package controller

import (
	"encoding/json"
	"net/http"
	"url-shortener/config"
	"url-shortener/model"
	"url-shortener/service"

	"gorm.io/gorm"
)

type UrlController struct {
	Service service.UrlService
	DB      *gorm.DB
}

func (ctrl *UrlController) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url, created, err := ctrl.Service.CreateShortUrl(req)
	if err != nil {
		http.Error(w, err.Error(), err.Code)
		return
	}

	res := model.CreateUrlResponse{
		Short:    url.Short,
		ShortUrl: config.BaseUrl + "/" + url.Short,
	}

	w.Header().Set("Content-Type", "application/json")
	if created {
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (ctrl *UrlController) DeleteShortUrl(w http.ResponseWriter, r *http.Request) {
	var req model.DeleteUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ctrl.Service.DeleteShortUrl(req)
	if err != nil {
		http.Error(w, err.Error(), err.Code)
		return
	}

	res := map[string]string{
		"message": "successfully deleted URL",
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
