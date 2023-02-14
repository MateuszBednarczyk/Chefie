package db

import (
	"back/src/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

var DB *gorm.DB

func Init(config *Config) {
	var err error
	dsn := "host=" + config.DbHost + " user=" + config.DbUsername + " dbname=" + config.DbName + " port=" + config.DbPort
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to db")
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Couldn't migrate")
	}
}
