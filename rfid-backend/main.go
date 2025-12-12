package main

import (
	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://rfid-system-1.onrender.com",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))
	docs.SwaggerInfo.BasePath = "/"

	api := r.Group("/api")
	{
		user_r := api.Group("/users")
		user_r.GET("/", controllers.GetUsers)
		user_r.POST("/", controllers.PostUser)
		user_r.PUT("/", controllers.PutUser)
		user_r.PUT("/:id", controllers.PutUserByID)
		user_r.GET("/card/:cardId", controllers.GetUserByCardID)
		user_r.PUT("/:id/classes", controllers.UpdateUserClasses)

		class_r := api.Group("/classes")
		class_r.GET("/", controllers.GetClasses)
		class_r.GET("/today", controllers.GetTodaysClasses)
		class_r.POST("/", controllers.PostClass)
		class_r.PUT("/:id", controllers.PutClassByID)
		class_r.DELETE("/:id", controllers.DeleteClass)

		attendence_r := api.Group("/attendance") 
		attendence_r.POST("/", controllers.PostAttendance)
		attendence_r.GET("/report/:classId", controllers.GetAttendanceReport)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
