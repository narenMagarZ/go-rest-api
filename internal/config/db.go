package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s port=%d password=%s dbname=%s",
		AppConfig().Db.DbHost,
		AppConfig().Db.DbUser,
		AppConfig().Db.DbPort,
		AppConfig().Db.DbPassword,
		AppConfig().Db.DbName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		panic(err)
	}
	return db;
}
