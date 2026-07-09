package service

import (
	"log"
	"net/http"
	"time"
	"url-shortener/config"
	"url-shortener/model"
	"url-shortener/persistence"
	"url-shortener/util"
)

type UrlService struct{}

func (s *UrlService) CreateShortUrl(req model.CreateUrlRequest) (model.URL, bool, *model.AppError) {
	var appErr model.AppError
	var err error
	url := model.URL{}

	result := persistence.DB.Where("original = ?", req.Url).Find(&url)
	if err = result.Error; err != nil {
		appErr = model.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		return url, false, &appErr
	}

	if result.RowsAffected > 0 {
		return url, false, nil
	}

	url.Original = req.Url
	url.Short, err = util.GenerateShortURL(8)
	if err != nil {
		appErr = model.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		return url, false, &appErr
	}

	var lifespan uint
	if req.Lifespan != nil {
		lifespan = *req.Lifespan
	} else {
		lifespan = config.DefaultLifespan
	}

	url.Expires = time.Now().Add(time.Duration(lifespan) * time.Minute)

	if err = persistence.DB.Create(&url).Error; err != nil {
		appErr = model.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		return url, false, &appErr
	}

	return url, true, nil
}

func (s *UrlService) DeleteShortUrl(req model.DeleteUrlRequest) *model.AppError {
	var appErr model.AppError
	var err error

	url := model.URL{}
	result := persistence.DB.Where("short = ?", req.Short).Delete(&url)
	if err = result.Error; err != nil {
		appErr = model.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		return &appErr
	}

	if result.RowsAffected == 0 {
		appErr = model.AppError{Code: http.StatusNotFound, Message: "URL not found"}
		return &appErr
	}

	return nil
}

func (s *UrlService) CheckForExpired() []model.URL {
	var urls []model.URL

	now := time.Now()
	result := persistence.DB.Where("expires <= ?", now).Find(&urls)
	if err := result.Error; err != nil {
		log.Println("something went wrong while searching for expired urls: ", err.Error())
		return urls
	}

	return urls
}
