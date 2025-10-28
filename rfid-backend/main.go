package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sonephyo/RFID_System/rfid-backend/controllers"
	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	user_r := r.Group("/users")
	user_r.GET("/", controllers.GetUser)

	// TODO: Implement attendence functionality after scanning
	// attendence_r := r.Group("/attendences")

	r.Run()
}
