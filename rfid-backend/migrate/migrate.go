package main

import (
	"log"

	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error: Migration Failed")
	}

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

	for _, dummy := range dummyData {
		result := initializers.DB.Create(&dummy)

		if result.Error != nil {
			log.Fatal("Error: error in adding dummy data into the db")
			return
		}
	}

}
