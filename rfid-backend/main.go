package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sonephyo/RFID_System/rfid-backend/controllers"
	docs "github.com/sonephyo/RFID_System/rfid-backend/docs"
	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	user_r := r.Group("/users")
	user_r.GET("/", controllers.GetUsers)
	user_r.POST("/", controllers.PostUser)
	user_r.PUT("/", controllers.PutUser)

	// TODO: Implement attendence functionality after scanning
	// attendence_r := r.Group("/attendences")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
