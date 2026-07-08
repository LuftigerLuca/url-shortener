package util

import (
	"log"
	"math/rand"
	"url-shortener/persistence"
)

func GenerateShortURL(length uint) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)

	for i := 0; i < int(length); i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}

	short := string(result)

	var exists bool
	if err := persistence.DB.Table("urls").Select("count(*) > 0").Where("short = ?", short).Find(&exists).Error; err != nil {
		log.Fatal("could not if short exists: ", err)
	}

	if exists {
		short = GenerateShortURL(length)
	}

	return short
}
