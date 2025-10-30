package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
	"gorm.io/gorm"
)

// @Summary get all available users from database
// @Tags users
// @Router /users [get]
func GetUsers(ctx *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"data": users,
	})
}

// @Summary create a new user
// @Schemes
// @Description create a new user if the user does not exist. If user does exist, it will return the already created user without returning anything
// @Tags users
// @Router /users [post]
func PostUser(ctx *gin.Context) {
	// Get data off req body
	var body struct {
		Name   string `json:"name"`
		Age    uint8  `json:"age"`
		CardID string `json:"cardID"`
	}

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	result := initializers.DB.Where(&models.User{Name: body.Name}).First(&existingUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			user := models.User{Name: body.Name, Age: body.Age, CardID: body.CardID}
			result = initializers.DB.Create(&user)

			if result.Error != nil {
				ctx.Status(400)
				return
			}

			ctx.JSON(200, gin.H{
				"user":    user,
				"message": "User created",
			})
			return
		} else {
			ctx.Status(400)
			return
		}
	}

	ctx.JSON(409, gin.H{
		"user":    existingUser,
		"message": "User already created",
	})
}
