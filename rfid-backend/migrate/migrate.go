package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {

	err := initializers.DB.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatalf("failed to drop table: %v", err)
	}
	fmt.Printf("Dropping User table...\n")

	err = initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error: Migration Failed")
	}
	fmt.Printf("User table created...\n")

	var dummyData = []models.User{
		{
			Name:   "Soney",
			Age:    21,
			CardID: "8061231234",
		},
		{
			Name:   "Vandan",
			Age:    22,
			CardID: "8064324321",
		},
	}

	for _, user := range dummyData {

		var existingUser models.User
		result := initializers.DB.Where(&models.User{Name: user.Name}).First(&existingUser)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				result := initializers.DB.Create(&user)
				if result.Error != nil {
					log.Fatal("Error: error in adding user data into the db")
					return
				}
				fmt.Printf("Added %s into User table\n", user.Name)
				continue
			} else {
				fmt.Printf("Error: %v\n", err)
				log.Fatal()
			}
		}

		fmt.Printf("User already exists: %s\n", existingUser.Name)
	}

}
