package models

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName string `gorm:"not null"`
	IsAdmin bool `gorm:"default:false"`
	Email string `gorm:"unique_index"` 
	Password string `gorm:"not null"`
}

type UserLogin struct{
	Email string
	Password string
}
