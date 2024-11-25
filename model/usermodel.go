package model

import "gorm.io/gorm"

type Usermodel struct {
	gorm.Model
	Email string `gorm:"not null"`
	Name  string `gorm:"not null"`
	Pass  string `gorm:"not null"`
}
