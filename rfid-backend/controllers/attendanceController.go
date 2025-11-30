package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
)

type AttendanceBody struct {
	CardID  string `json:"cardId"`
	ClassID uint   `json:"classId"`
}

func PostAttendance(ctx *gin.Context) {
	var body AttendanceBody
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Find user by card
	var user models.User
	if err := initializers.DB.Preload("Classes").Where("card_id = ?", body.CardID).First(&user).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Card not registered", "success": false})
		return
	}

	// Check if user is enrolled in this class
	enrolled := false
	for _, class := range user.Classes {
		if class.ID == body.ClassID {
			enrolled = true
			break
		}
	}
	if !enrolled {
		ctx.JSON(403, gin.H{"error": "Not enrolled in class", "success": false, "name": user.Name})
		return
	}

	// Check if already checked in today for this class
	today := time.Now().Format("2006-01-02")
	var existing models.Attendance
	err := initializers.DB.Where("user_id = ? AND class_id = ? AND DATE(check_in_time) = ?", user.ID, body.ClassID, today).First(&existing).Error
	if err == nil {
		ctx.JSON(409, gin.H{"error": "Already checked in", "success": false, "name": user.Name})
		return
	}

	// Record attendance
	attendance := models.Attendance{
		UserID:      user.ID,
		ClassID:     body.ClassID,
		CheckInTime: time.Now(),
	}
	initializers.DB.Create(&attendance)

	ctx.JSON(200, gin.H{"success": true, "name": user.Name, "message": "Checked in"})
}

func GetAttendanceByClass(ctx *gin.Context) {
	classId := ctx.Param("classId")

	var attendances []models.Attendance
	initializers.DB.Preload("User").Where("class_id = ?", classId).Find(&attendances)

	ctx.JSON(200, gin.H{"data": attendances})
}