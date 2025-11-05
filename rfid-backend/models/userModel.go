package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  uint8
	CardID string
	Classes []*Class `gorm:"many2many:user_classes;"`
}
