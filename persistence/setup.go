package persistence

import (
	"fmt"
	"log"
	"url-shortener/config"
	"url-shortener/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	DBURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	DB, err = gorm.Open(mysql.Open(DBURI), &gorm.Config{})
	if err != nil {
		log.Fatal("connection error:", err)
	}

	log.Println("we are connected to the database", config.DBHost)
	if err := DB.AutoMigrate(&model.URL{}); err != nil {
		log.Fatal("AutoMigrate error:", err)
	}
}
