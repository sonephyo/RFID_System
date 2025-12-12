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

func GetAttendanceReport(ctx *gin.Context) {
	classId := ctx.Param("classId")
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")

	var class models.Class
	if err := initializers.DB.Preload("Users").First(&class, classId).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Class not found"})
		return
	}

	query := initializers.DB.Where("class_id = ?", classId)
	if startDate != "" {
		query = query.Where("check_in_time >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("check_in_time <= ?", endDate+" 23:59:59")
	}

	var attendances []models.Attendance
	query.Find(&attendances)

	userAttendance := make(map[uint]int)
	for _, a := range attendances {
		userAttendance[a.UserID]++
	}

	dateSet := make(map[string]bool)
	for _, a := range attendances {
		dateStr := a.CheckInTime.Format("2006-01-02")
		dateSet[dateStr] = true
	}
	totalSessions := len(dateSet)

	if totalSessions == 0 {
		totalSessions = 1
	}

	type UserStats struct {
		ID             uint    `json:"id"`
		Name           string  `json:"name"`
		CardID         string  `json:"cardId"`
		AttendedCount  int     `json:"attendedCount"`
		TotalSessions  int     `json:"totalSessions"`
		AttendanceRate float64 `json:"attendanceRate"`
	}

	var userStats []UserStats
	for _, user := range class.Users {
		attended := userAttendance[user.ID]
		rate := float64(attended) / float64(totalSessions) * 100
		if rate > 100 {
			rate = 100
		}
		userStats = append(userStats, UserStats{
			ID:             user.ID,
			Name:           user.Name,
			CardID:         user.CardID,
			AttendedCount:  attended,
			TotalSessions:  totalSessions,
			AttendanceRate: rate,
		})
	}

	var totalRate float64
	for _, u := range userStats {
		totalRate += u.AttendanceRate
	}
	classAverage := 0.0
	if len(userStats) > 0 {
		classAverage = totalRate / float64(len(userStats))
	}

	ctx.JSON(200, gin.H{
		"class": gin.H{
			"id":            class.ID,
			"name":          class.Name,
			"totalSessions": totalSessions,
			"averageRate":   classAverage,
			"enrolledCount": len(class.Users),
		},
		"users": userStats,
	})
}
