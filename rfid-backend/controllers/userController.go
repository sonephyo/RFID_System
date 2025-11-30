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
	result := initializers.DB.Preload("Classes").Find(&users)
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
	var body UserBody
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if card already registered
	var existingUser models.User
	result := initializers.DB.Where("card_id = ?", body.CardID).First(&existingUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			user := models.User{Name: body.Name, Age: body.Age, CardID: body.CardID}
			result = initializers.DB.Create(&user)
			if result.Error != nil {
				ctx.JSON(400, gin.H{"error": "Failed to create user"})
				return
			}
			ctx.JSON(200, gin.H{
				"user":    user,
				"message": "User created",
			})
			return
		} else {
			ctx.JSON(400, gin.H{"error": "Database error"})
			return
		}
	}
	ctx.JSON(409, gin.H{
		"error":   "Card already registered",
		"user":    existingUser,
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

// @Summary Get user by Card ID
// @Tags users
// @Produce json
// @Param cardId path string true "Card ID"
// @Router /users/card/{cardId} [get]
func GetUserByCardID(ctx *gin.Context) {
	cardID := ctx.Param("cardId")

	var user models.User
	result := initializers.DB.Preload("Classes").Where("card_id = ?", cardID).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(200, user)
}

// @Summary Update user's classes
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Router /users/{id}/classes [put]
func UpdateUserClasses(ctx *gin.Context) {
	userID := ctx.Param("id")

	var body struct {
		ClassIDs []uint `json:"classIds"`
	}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var classes []models.Class
	initializers.DB.Find(&classes, body.ClassIDs)

	initializers.DB.Model(&user).Association("Classes").Replace(classes)

	initializers.DB.Preload("Classes").First(&user, userID)
	ctx.JSON(200, user)
}

// @Summary Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Router /users/{id} [put]
func PutUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	var body struct {
		Name string `json:"Name"`
		Age  uint8  `json:"Age"`
	}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.Name = body.Name
	user.Age = body.Age
	initializers.DB.Save(&user)

	initializers.DB.Preload("Classes").First(&user, userID)
	ctx.JSON(200, user)
}