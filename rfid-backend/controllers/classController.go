package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sonephyo/RFID_System/rfid-backend/initializers"
	"github.com/sonephyo/RFID_System/rfid-backend/models"
)

type ClassBody struct {
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Monday    bool   `json:"monday"`
	Tuesday   bool   `json:"tuesday"`
	Wednesday bool   `json:"wednesday"`
	Thursday  bool   `json:"thursday"`
	Friday    bool   `json:"friday"`
	Saturday  bool   `json:"saturday"`
	Sunday    bool   `json:"sunday"`
}

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

// @Summary Get today's classes
// @Tags classes
// @Produce json
// @Router /classes/today [get]
func GetTodaysClasses(ctx *gin.Context) {
	dayMap := map[string]string{
		"Monday":    "monday",
		"Tuesday":   "tuesday",
		"Wednesday": "wednesday",
		"Thursday":  "thursday",
		"Friday":    "friday",
		"Saturday":  "saturday",
		"Sunday":    "sunday",
	}

	today := time.Now().Weekday().String()
	column := dayMap[today]

	var classes []models.Class
	initializers.DB.Where(column+" = ?", true).Find(&classes)

	ctx.JSON(200, gin.H{"data": classes})
}

// @Summary Create a new class
// @Tags classes
// @Accept json
// @Produce json
// @Router /classes [post]
func PostClass(ctx *gin.Context) {
	var body ClassBody
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	class := models.Class{
		Name:      body.Name,
		StartTime: body.StartTime,
		EndTime:   body.EndTime,
		Monday:    body.Monday,
		Tuesday:   body.Tuesday,
		Wednesday: body.Wednesday,
		Thursday:  body.Thursday,
		Friday:    body.Friday,
		Saturday:  body.Saturday,
		Sunday:    body.Sunday,
	}

	result := initializers.DB.Create(&class)
	if result.Error != nil {
		ctx.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(200, class)
}

// @Summary Update class by ID
// @Tags classes
// @Accept json
// @Produce json
// @Router /classes/{id} [put]
func PutClassByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var class models.Class
	if err := initializers.DB.First(&class, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Class not found"})
		return
	}

	var body ClassBody
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&class).Updates(models.Class{
		Name:      body.Name,
		StartTime: body.StartTime,
		EndTime:   body.EndTime,
		Monday:    body.Monday,
		Tuesday:   body.Tuesday,
		Wednesday: body.Wednesday,
		Thursday:  body.Thursday,
		Friday:    body.Friday,
		Saturday:  body.Saturday,
		Sunday:    body.Sunday,
	})

	ctx.JSON(200, class)
}

// @Summary Delete class by ID
// @Tags classes
// @Produce json
// @Router /classes/{id} [delete]
func DeleteClass(ctx *gin.Context) {
	id := ctx.Param("id")

	var class models.Class
	if err := initializers.DB.First(&class, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Class not found"})
		return
	}

	initializers.DB.Delete(&class)
	ctx.JSON(200, gin.H{"message": "Class deleted"})
}