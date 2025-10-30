package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
	"gorm.io/gorm"
)

type UserBody struct {
	Name   string `json:"name"`
	Age    uint8  `json:"age"`
	CardID string `json:"cardID"`
}

// @Summary get all available users from database
// @Tags users
// @Produce json
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

// @Summary Create a new user
// @Description This endpoint creates a new user.
// If the user already exists, it returns the existing user instead of creating a new one.
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserBody true "UserBody data"
// @Success 200 {object} UserBody "Successfully created or returned existing user"
// @Router /users [post]
func PostUser(ctx *gin.Context) {
	// Get data off req body
	var body UserBody

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


// @Summary Update existing user
// @Description This endpoint updates existing user.
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserBody true "UserBody data"
// @Success 200 {object} UserBody "Successfully updated user"
// @Router /users [put]
func PutUser(ctx *gin.Context) {
	var body UserBody

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	result := initializers.DB.Where(&models.User{Name: body.Name}).First(&existingUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.Status(404)
		} else {
			ctx.Status(400)
		}
		return
	}

	existingUser.Age = body.Age
	existingUser.CardID = body.CardID

	result = initializers.DB.Save(&existingUser)

	if result.Error != nil {
		ctx.Status(400)
		return
	}
	
	ctx.JSON(200, gin.H{
		"user":    existingUser,
		"message": "User updated",
	})

}
