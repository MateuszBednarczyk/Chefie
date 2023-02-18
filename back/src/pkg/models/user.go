package models

type User struct {
	ID           int64  `gorm:"primary_key;auto_increment;uniqueIndex;not null'" json:"ID"`
	Username     string `gorm:"uniqueIndex;notnull"`
	PasswordHash string `gorm:"notnull"`
}
