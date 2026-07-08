package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DefaultLifespan int
var BaseUrl string
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

	DefaultLifespan, err = strconv.Atoi(os.Getenv("DEFAULT_LIFESPAN"))
	if err != nil {
		log.Fatal("cannot parse default lifespan:", err)
	}

	BaseUrl = os.Getenv("BASE_URL")

	DBDriver = os.Getenv("DB_DRIVER")
	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")
}
