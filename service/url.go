package service

import (
	"net/http"
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

	if req.Lifespan != nil {
		url.Lifespan = *req.Lifespan
	} else {
		url.Lifespan = config.DefaultLifespan
	}

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
