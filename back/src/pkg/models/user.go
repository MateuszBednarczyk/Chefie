package models

type User struct {
	ID           uint `gorm:"primary_key;auto_increment;uniqueIndex;not null'" json:"ID"`
	Username     string
	PasswordHash string
}
