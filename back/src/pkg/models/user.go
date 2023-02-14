package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint `gorm:"primary_key;auto_increment" json:"ID"`
	Username     string
	PasswordHash string
}
