package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DefaultLifespan uint
var BaseUrl string
var CleanupInterval uint
var DBDriver string
var DBHost string
var DBUser string
var DBPassword string
var DBName string
var DBPort string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file:", err)
	}

	lifespan, err := strconv.ParseUint(os.Getenv("DEFAULT_LIFESPAN"), 10, 64)
	if err != nil {
		log.Fatal("cannot parse default lifespan:", err)
	}
	DefaultLifespan = uint(lifespan)

	cleanup, err := strconv.ParseUint(os.Getenv("CLEANUP_INTERVAL"), 10, 64)
	if err != nil {
		log.Fatal("cannot parse cleanup interval:", err)
	}
	CleanupInterval = uint(cleanup)

	BaseUrl = os.Getenv("BASE_URL")

	DBDriver = os.Getenv("DB_DRIVER")
	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")
}
