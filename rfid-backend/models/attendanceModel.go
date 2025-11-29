// models/attendance.go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Attendance struct {
    gorm.Model
    UserID      uint
    ClassID     uint
    CheckInTime time.Time
}