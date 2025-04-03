package config

import (
	"fmt"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	username := helper.GoDotEnv("DB_USERNAME")
	password := helper.GoDotEnv("DB_PASSWORD")
	name := helper.GoDotEnv("DB_NAME")
	port := helper.GoDotEnv("DB_PORT")
	host := helper.GoDotEnv("DB_HOST")
	timezone := helper.GoDotEnv("DB_TIMEZONE")

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
		PrepareStmt:            true,
		SkipDefaultTransaction: false,
	})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Member{},
		&model.Religion{},
		&model.RWProfile{},
		&model.Village{},
		&model.ReferalCode{},
		&model.Religion{},
		&model.RWLeader{},
		&model.RTLeader{},
		&model.RTProfile{},
		&model.IncomingLetter{},
		&model.OutcomingLetter{},
		&model.RegistrationToken{},
		&model.LetterType{},
		&model.Inventory{},
		&model.Income{},
		&model.IncomePlan{},
		&model.Expense{},
	)

	return nil
}
