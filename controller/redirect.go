package controller

import (
	"net/http"
	"time"
	"url-shortener/model"

	"gorm.io/gorm"
)

type RedirectController struct {
	DB *gorm.DB
}

func (ctrl *RedirectController) Redirect(w http.ResponseWriter, r *http.Request) {
	short := r.PathValue("short")

	url := model.URL{}
	result := ctrl.DB.Where("short = ?", short).Find(&url)
	if err := result.Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if url.Original == "" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	if time.Now().After(url.Expires) {
		http.Error(w, "URL has expired", http.StatusGone)
		return
	}

	http.Redirect(w, r, url.Original, http.StatusMovedPermanently)
}
