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
		{Name: "CSC322", StartTime: "09:00", EndTime: "10:30", Monday: true, Wednesday: true, Friday: true},
		{Name: "CSC473", StartTime: "14:00", EndTime: "15:30", Tuesday: true, Thursday: true},
		{Name: "CSC101", StartTime: "11:00", EndTime: "12:00", Monday: true, Wednesday: true},
	}
	initializers.DB.Create(&classData)
	fmt.Println("Created 3 classes...")

	userData := []models.User{
		{
			Name:   "Soney",
			Age:    21,
			CardID: "8061231234",
			Classes: []*models.Class{
				classData[0],
				classData[1],
			},
		},
		{
			Name:   "Vandan",
			Age:    22,
			CardID: "8064324321",
			Classes: []*models.Class{
				classData[0],
			},
		},
		{
			Name:   "Alice",
			Age:    20,
			CardID: "1234567890",
			Classes: []*models.Class{
				classData[0],
				classData[1],
				classData[2],
			},
		},
		{
			Name:   "Bob",
			Age:    23,
			CardID: "0987654321",
			Classes: []*models.Class{
				classData[2],
			},
		},
	}
	initializers.DB.Create(&userData)
	fmt.Println("Created 4 users...")

	now := time.Now()
	attendanceData := []models.Attendance{
		{UserID: userData[0].ID, ClassID: classData[0].ID, CheckInTime: now.Add(-24 * time.Hour)},
		{UserID: userData[0].ID, ClassID: classData[1].ID, CheckInTime: now.Add(-1 * time.Hour)},
		{UserID: userData[1].ID, ClassID: classData[0].ID, CheckInTime: now.Add(-24 * time.Hour)},
		{UserID: userData[2].ID, ClassID: classData[0].ID, CheckInTime: now.Add(-2 * time.Hour)},
		{UserID: userData[2].ID, ClassID: classData[2].ID, CheckInTime: now.Add(-30 * time.Minute)},
	}
	initializers.DB.Create(&attendanceData)
	fmt.Println("Created 5 attendance records...")

	fmt.Println("\n=== Seed Complete ===")
	fmt.Println("Users: Soney (8061231234), Vandan (8064324321), Alice (1234567890), Bob (0987654321)")
	fmt.Println("Classes: CSC322, CSC473, CSC101")
}
