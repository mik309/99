package models

import (
	"gorm.io/gorm"
)

type Address struct{
	gorm.Model 
	Coordinates string  `gorm:"type:varchar(100);not null"`
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName string `gorm:"type:varchar(100);not null"`
	Street string `gorm:"not null"`
	ZipCode string `gorm:"type:varchar(10);not null"`
	State string `gorm:"type:varchar(100);not null"`
	City string `gorm:"type:varchar(100);not null"`
	Neighbourhood string `gorm:"type:varchar(100);not null"`
	ExNumber string `gorm:"type:varchar(20);not null"`
	InNumber string `gorm:"type:varchar(20);"`
	PhoneNumber string `gorm:"type:varchar(20);not null"`
	OrderID uint 
}

type Product struct{
	gorm.Model
	Weight float64 `gorm:"not null"`
	OrderID uint
}


type Order struct{
	gorm.Model
	PackageSize string `gorm:"not null"`
	Status string `gorm:"not null;default:creado"`
	Refund bool `gorm:"default:false"`
	DestinationAddress Address `gorm:"foreignKey:OrderID"`
	Products []Product `gorm:"foreignKey:OrderID"`
}