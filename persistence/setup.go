package persistence

import (
	"fmt"
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
		slog.Error("database connection error:", "error", err.Error())
	}

	slog.Info("connected to database at", "host", config.DBHost)
	if err := DB.AutoMigrate(&model.URL{}); err != nil {
		slog.Error("AutoMigrate error:", "error", err.Error())
	}

	return DB
}
