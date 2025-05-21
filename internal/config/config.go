package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	DbHost string
	DbUser string
	DbPassword string
	DbName string
	DbPort int
}

type appConfig struct {
	JwtSecretKey string
	Port string
	Db dbConfig
}

func AppConfig() appConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT")) // convert ascii to int
	if err != nil {
		log.Fatal("Failed to get port ", err);
	}
	
	appConfigData := appConfig{ 
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
		Port: os.Getenv("PORT"),
		Db: dbConfig{
			DbHost: os.Getenv("DB_HOST"),
			DbUser: os.Getenv("DB_USER"),
			DbPort: port,
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName: os.Getenv("DB_NAME"),
		},
	}

	return appConfigData

}