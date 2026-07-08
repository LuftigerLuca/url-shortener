package util

import (
	"math/rand"
	"url-shortener/persistence"
)

func GenerateShortURL(length uint) (string, error) {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)

	for i := 0; i < int(length); i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}

	var err error
	short := string(result)

	var exists bool
	if err = persistence.DB.Table("urls").Select("count(*) > 0").Where("short = ?", short).Find(&exists).Error; err != nil {
		return "", err
	}

	if exists {
		short, err = GenerateShortURL(length)
	}

	return short, err
}
