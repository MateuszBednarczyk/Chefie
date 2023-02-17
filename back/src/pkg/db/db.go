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

var db *gorm.DB

func Init(config *Config) {
	var err error
	dsn := "host=" + config.DbHost + " user=" + config.DbUsername + " dbname=" + config.DbName + " port=" + config.DbPort
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to db")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Couldn't migrate")
	}
}

func GetDb() *gorm.DB {
	return db
}
