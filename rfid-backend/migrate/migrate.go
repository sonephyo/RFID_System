package main

import (
	"fmt"
	"log"

	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {

	err := initializers.DB.Migrator().DropTable(&models.User{}, &models.Class{}, "user_classes")
	if err != nil {
		log.Fatalf("failed to drop tables: %v", err)
	}
	fmt.Printf("Dropping tables...\n")

	err = initializers.DB.AutoMigrate(&models.User{}, &models.Class{})
	if err != nil {
		log.Fatal("Error: Migration Failed")
	}
	fmt.Printf("Tables created...\n")

	classData := []*models.Class{
		{Name: "CSC322"},
		{Name: "CSC473"},
	}
	initializers.DB.Create(&classData)

	var userData = []models.User{
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
	}

	initializers.DB.Create(userData)
}
