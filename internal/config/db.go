package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading env", err);
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT")) // convert ascii to int
	if err != nil {
		log.Fatal("Failed to get port ", err);
	}
	dsn := fmt.Sprintf("host=%s user=%s port=%d password=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		port,
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		panic(err)
	}
	return db;
}
