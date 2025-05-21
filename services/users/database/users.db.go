package databaseusers

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;auto_increment"`
	FirstName string
	LastName  string
	Nickname  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	Country   int `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Countries struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;auto_increment"`
	Code  string `gorm:"unique"`
	Name  string
	Users []Users `gorm:"foreignKey:Country"`
}
