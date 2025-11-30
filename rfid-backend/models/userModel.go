// models/user.go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name    string
    Age     uint8
    CardID  string   `gorm:"unique"`
    Classes []*Class `gorm:"many2many:user_classes;"`
}