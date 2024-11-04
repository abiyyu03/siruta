package config

import (
	"fmt"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	username := usecase.GoDotEnv("DB_USERNAME")
	password := usecase.GoDotEnv("DB_PASSWORD")
	name := usecase.GoDotEnv("DB_NAME")
	port := usecase.GoDotEnv("DB_PORT")
	host := usecase.GoDotEnv("DB_HOST")
	timezone := usecase.GoDotEnv("DB_TIMEZONE")

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host,
		username,
		password,
		name,
		port,
		timezone,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&model.User{},
		&model.Role{},
	)

	return nil
}
