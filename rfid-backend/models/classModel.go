// models/class.go
package models

import "gorm.io/gorm"

type Class struct {
    gorm.Model
    Name  string
    Users []*User `gorm:"many2many:user_classes;"`
}