package db

import (
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

func Init(config *Config) {
	dsn := "host=" + config.DbHost + " user=" + config.DbUsername + " dbname=" + config.DbName + " port=" + config.DbPort
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to db")
	}
	err = db.AutoMigrate()
	if err != nil {
		panic("Couldn't migrate")
	}
}
