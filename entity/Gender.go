package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Name string

	User []User `gorm:"foreignkey:GenderID"`
}