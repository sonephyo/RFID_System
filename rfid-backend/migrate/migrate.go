package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.Migrator().DropTable(
		&models.Attendance{},
		&models.User{},
		&models.Class{},
		"user_classes",
	)
	if err != nil {
		log.Fatalf("failed to drop tables: %v", err)
	}
	fmt.Println("Dropped tables...")

	err = initializers.DB.AutoMigrate(&models.User{}, &models.Class{}, &models.Attendance{})
	if err != nil {
		log.Fatal("Error: Migration Failed")
	}
	fmt.Println("Tables created...")

	classData := []*models.Class{
		{Name: "CSC322", StartTime: "11:30", EndTime: "12:25", Monday: true, Wednesday: true, Friday: true},
		{Name: "CSC473", StartTime: "09:10", EndTime: "10:05", Monday: true, Wednesday: true, Friday: true},
		{Name: "CSC496", StartTime: "10:20", EndTime: "11:10", Monday: true, Wednesday: true, Friday: true},
		{Name: "CSC461", StartTime: "9:35", EndTime: "10:55", Tuesday: true, Thursday: true},
	}
	initializers.DB.Create(&classData)
	fmt.Println("Created 3 classes...")

	userData := []models.User{
		{
			Name:   "Soney",
			Age:    21,
			CardID: "04c37982c61190",
			Classes: []*models.Class{
				classData[0],
				classData[1],
				classData[2],
				classData[3],
			},
		},
		{
			Name:   "Vandan",
			Age:    22,
			CardID: "04b6c182c61190",
			Classes: []*models.Class{
				classData[0],
				classData[1],
			},
		},
		{
			Name:    "Dummy1",
			Age:     20,
			CardID:  "0430e632c61191",
			Classes: []*models.Class{},
		},
	}
	initializers.DB.Create(&userData)
	fmt.Println("Created 4 users...")

	attendanceData := []models.Attendance{
		// ========================================
		// SONEY (userData[0]) - Enrolled in CSC322, CSC473, CSC496, CSC461
		// Good student but skips more as semester progresses, especially Fridays
		// ========================================

		// CSC473 (MWF 9:10) - Soney
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 8, 25, 9, 8, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 8, 27, 9, 7, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 8, 29, 9, 9, 0, 0, time.Local)},
		// Sep 1 - Labor Day
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 3, 9, 8, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 5, 9, 10, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 8, 9, 7, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 10, 9, 9, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 12, 9, 8, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 15, 9, 6, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 17, 9, 11, 0, 0, time.Local)},
		// Sep 19 - missed (Friday)
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 22, 9, 9, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 24, 9, 7, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 26, 9, 10, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 29, 9, 8, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 1, 9, 9, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 3, 9, 7, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 6, 9, 8, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 8, 9, 10, 0, 0, time.Local)},
		// Oct 10 - missed (Friday)
		// Oct 13-14 Fall Break
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 15, 9, 8, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 17, 9, 7, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 20, 9, 9, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 22, 9, 8, 0, 0, time.Local)},
		// Oct 24 - missed (Friday)
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 27, 9, 7, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 29, 9, 22, 0, 0, time.Local)}, // very late
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 31, 9, 8, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 11, 3, 9, 6, 0, 0, time.Local)},
		// Nov 5 - missed
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 11, 7, 9, 8, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 11, 10, 9, 10, 0, 0, time.Local)},
		// Nov 12 - missed
		// Nov 14 - missed (burnout week)
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 11, 17, 9, 15, 0, 0, time.Local)}, // late
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 11, 19, 9, 7, 0, 0, time.Local)},
		// Nov 21 - missed
		// Nov 24-28 Thanksgiving Break
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 12, 1, 9, 8, 0, 0, time.Local)},
		// Dec 3 - missed (finals prep)
		// Dec 5 - missed (finals prep)

		// CSC496 (MWF 10:20) - Soney
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 8, 25, 10, 18, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 8, 27, 10, 19, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 8, 29, 10, 17, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 3, 10, 20, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 5, 10, 18, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 8, 10, 19, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 10, 10, 21, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 12, 10, 18, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 15, 10, 17, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 17, 10, 20, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 19, 10, 19, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 22, 10, 18, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 24, 10, 22, 0, 0, time.Local)},
		// Sep 26 - missed
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 9, 29, 10, 19, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 1, 10, 18, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 3, 10, 20, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 6, 10, 17, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 8, 10, 19, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 10, 10, 21, 0, 0, time.Local)},
		// Oct 13-14 Fall Break
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 15, 10, 18, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 17, 10, 20, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 20, 10, 19, 0, 0, time.Local)},
		// Oct 22 - missed
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 24, 10, 18, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 27, 10, 17, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 10, 29, 10, 20, 0, 0, time.Local)},
		// Oct 31 - missed (Halloween)
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 11, 3, 10, 19, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 11, 5, 10, 18, 0, 0, time.Local)},
		// Nov 7 - missed
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 11, 10, 10, 32, 0, 0, time.Local)}, // very late
		// Nov 12 - missed
		// Nov 14 - missed
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 11, 17, 10, 20, 0, 0, time.Local)},
		// Nov 19 - missed
		// Nov 21 - missed
		// Nov 24-28 Thanksgiving Break
		// Dec 1 - missed (finals crunch)
		{UserID: userData[0].ID, ClassID: classData[2].ID, CheckInTime: time.Date(2025, 12, 3, 10, 25, 0, 0, time.Local)}, // late
		// Dec 5 - missed

		// CSC322 (MWF 11:30) - Soney
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 8, 25, 11, 28, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 8, 27, 11, 29, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 8, 29, 11, 27, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 3, 11, 30, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 5, 11, 28, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 8, 11, 29, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 10, 11, 31, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 12, 11, 28, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 15, 11, 27, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 17, 11, 30, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 19, 11, 29, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 22, 11, 28, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 24, 11, 32, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 26, 11, 29, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 29, 11, 30, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 1, 11, 28, 0, 0, time.Local)},
		// Oct 3 - missed
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 6, 11, 29, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 8, 11, 27, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 10, 11, 30, 0, 0, time.Local)},
		// Oct 13-14 Fall Break
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 15, 11, 28, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 17, 11, 31, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 20, 11, 29, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 22, 11, 28, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 24, 11, 30, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 27, 11, 27, 0, 0, time.Local)},
		// Oct 29 - missed
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 31, 11, 29, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 11, 3, 11, 28, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 11, 5, 11, 30, 0, 0, time.Local)},
		// Nov 7 - missed
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 11, 10, 11, 29, 0, 0, time.Local)},
		// Nov 12 - missed
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 11, 14, 11, 45, 0, 0, time.Local)}, // very late
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 11, 17, 11, 28, 0, 0, time.Local)},
		// Nov 19 - missed
		// Nov 21 - missed
		// Nov 24-28 Thanksgiving Break
		// Dec 1 - missed
		// Dec 3 - missed
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 12, 5, 11, 35, 0, 0, time.Local)}, // late

		// CSC461 (TTh 9:35) - Soney
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 8, 26, 9, 33, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 8, 28, 9, 32, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 9, 2, 9, 34, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 9, 4, 9, 31, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 9, 9, 9, 33, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 9, 11, 9, 35, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 9, 16, 9, 32, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 9, 18, 9, 34, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 9, 23, 9, 33, 0, 0, time.Local)},
		// Sep 25 - missed
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 9, 30, 9, 31, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 10, 2, 9, 34, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 10, 7, 9, 32, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 10, 9, 9, 33, 0, 0, time.Local)},
		// Oct 14 Fall Break
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 10, 16, 9, 35, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 10, 21, 9, 32, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 10, 23, 9, 34, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 10, 28, 9, 33, 0, 0, time.Local)},
		// Oct 30 - missed
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 11, 4, 9, 31, 0, 0, time.Local)},
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 11, 6, 9, 48, 0, 0, time.Local)}, // late
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 11, 11, 9, 33, 0, 0, time.Local)},
		// Nov 13 - missed
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 11, 18, 9, 34, 0, 0, time.Local)},
		// Nov 20 - missed
		// Nov 25-27 Thanksgiving
		// Dec 2 - missed
		{UserID: userData[0].ID, ClassID: classData[3].ID, CheckInTime: time.Date(2025, 12, 4, 9, 40, 0, 0, time.Local)}, // late

		// ========================================
		// VANDAN (userData[1]) - Enrolled in CSC322, CSC473
		// Decent attendance but misses more classes, often late
		// ========================================

		// CSC473 (MWF 9:10) - Vandan
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 8, 25, 9, 12, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 8, 27, 9, 15, 0, 0, time.Local)}, // late
		// Aug 29 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 3, 9, 9, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 5, 9, 11, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 8, 9, 18, 0, 0, time.Local)}, // late
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 10, 9, 10, 0, 0, time.Local)},
		// Sep 12 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 15, 9, 8, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 17, 9, 14, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 19, 9, 9, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 22, 9, 11, 0, 0, time.Local)},
		// Sep 24 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 26, 9, 10, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 9, 29, 9, 9, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 1, 9, 12, 0, 0, time.Local)},
		// Oct 3 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 6, 9, 8, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 8, 9, 16, 0, 0, time.Local)}, // late
		// Oct 10 - missed
		// Oct 13-14 Fall Break
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 15, 9, 11, 0, 0, time.Local)},
		// Oct 17 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 20, 9, 9, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 22, 9, 13, 0, 0, time.Local)},
		// Oct 24 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 27, 9, 10, 0, 0, time.Local)},
		// Oct 29 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 10, 31, 9, 22, 0, 0, time.Local)}, // very late
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 11, 3, 9, 9, 0, 0, time.Local)},
		// Nov 5 - missed
		// Nov 7 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 11, 10, 9, 11, 0, 0, time.Local)},
		// Nov 12 - missed
		// Nov 14 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 11, 17, 9, 8, 0, 0, time.Local)},
		// Nov 19 - missed
		// Nov 21 - missed
		// Nov 24-28 Thanksgiving Break
		// Dec 1 - missed
		{UserID: userData[1].ID, ClassID: classData[1].ID, CheckInTime: time.Date(2025, 12, 3, 9, 25, 0, 0, time.Local)}, // very late
		// Dec 5 - missed

		// CSC322 (MWF 11:30) - Vandan
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 8, 25, 11, 32, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 8, 27, 11, 29, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 8, 29, 11, 35, 0, 0, time.Local)}, // late
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 3, 11, 30, 0, 0, time.Local)},
		// Sep 5 - missed
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 8, 11, 31, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 10, 11, 28, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 12, 11, 33, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 15, 11, 29, 0, 0, time.Local)},
		// Sep 17 - missed
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 19, 11, 30, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 22, 11, 28, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 24, 11, 38, 0, 0, time.Local)}, // late
		// Sep 26 - missed
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 9, 29, 11, 31, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 1, 11, 29, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 3, 11, 30, 0, 0, time.Local)},
		// Oct 6 - missed
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 8, 11, 32, 0, 0, time.Local)},
		// Oct 10 - missed
		// Oct 13-14 Fall Break
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 15, 11, 29, 0, 0, time.Local)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 17, 11, 31, 0, 0, time.Local)},
		// Oct 20 - missed
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 22, 11, 30, 0, 0, time.Local)},
		// Oct 24 - missed
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 10, 27, 11, 28, 0, 0, time.Local)},
		// Oct 29 - missed
		// Oct 31 - missed
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 11, 3, 11, 42, 0, 0, time.Local)}, // very late

		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 11, 10, 11, 30, 0, 0, time.Local)},

		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: time.Date(2025, 11, 19, 11, 35, 0, 0, time.Local)}, // late
	}
	initializers.DB.Create(&attendanceData)
	fmt.Println("Created 5 attendance records...")

	fmt.Println("\n=== Seed Complete ===")
}
