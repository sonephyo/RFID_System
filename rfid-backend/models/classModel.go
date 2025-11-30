// models/class.go
package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name      string  `gorm:"not null" json:"name"`
	StartTime string  `gorm:"not null" json:"startTime"`
	EndTime   string  `gorm:"not null" json:"endTime"`
	Monday    bool    `gorm:"default:false" json:"monday"`
	Tuesday   bool    `gorm:"default:false" json:"tuesday"`
	Wednesday bool    `gorm:"default:false" json:"wednesday"`
	Thursday  bool    `gorm:"default:false" json:"thursday"`
	Friday    bool    `gorm:"default:false" json:"friday"`
	Saturday  bool    `gorm:"default:false" json:"saturday"`
	Sunday    bool    `gorm:"default:false" json:"sunday"`
	Users     []*User `gorm:"many2many:user_classes;" json:"users,omitempty"`
}