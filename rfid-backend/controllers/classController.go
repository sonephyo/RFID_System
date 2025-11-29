package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
)

// @Summary Get all classes
// @Tags classes
// @Produce json
// @Router /classes [get]
func GetClasses(ctx *gin.Context) {
	var classes []models.Class
	result := initializers.DB.Find(&classes)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"data": classes,
	})
}