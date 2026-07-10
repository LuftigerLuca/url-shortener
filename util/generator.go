package util

import (
	"math/rand"

	"gorm.io/gorm"
)

func GenerateShortURL(length uint, db *gorm.DB) (string, error) {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)

	for i := 0; i < int(length); i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}

	var err error
	short := string(result)

	var exists bool
	if err = db.Table("urls").Select("count(*) > 0").Where("short = ?", short).Find(&exists).Error; err != nil {
		return "", err
	}

	if exists {
		short, err = GenerateShortURL(length, db)
	}

	return short, err
}
