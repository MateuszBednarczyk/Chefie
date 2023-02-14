package models

type User struct {
	ID           uint `gorm:"primary_key;auto_increment;unique;notnull'" json:"ID"`
	Username     string
	PasswordHash string
}
