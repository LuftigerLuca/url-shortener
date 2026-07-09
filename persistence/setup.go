package persistence

import (
	"fmt"
	"log"
	"log/slog"
	"url-shortener/config"
	"url-shortener/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	DBURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	DB, err := gorm.Open(mysql.Open(DBURI), &gorm.Config{})
	if err != nil {
		slog.Error("database connection error:", err)
	}

	slog.Info("connected to database at", config.DBHost)
	if err := DB.AutoMigrate(&model.URL{}); err != nil {
		log.Fatal("AutoMigrate error:", err)
	}

	return DB
}
