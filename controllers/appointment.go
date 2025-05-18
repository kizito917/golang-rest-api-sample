package controllers

import (
	"golang-rest-api/config"
	"golang-rest-api/middlewares"
	"golang-rest-api/models"

	"github.com/gin-gonic/gin"
)

func CreateNewAppointment(c *gin.Context) {
	userId, err := middlewares.GetUserIDFromClaims(c)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "User details not found, Invalid user ID",
			"Error": err,
		})
		return
	}

	var user models.User
	var appointmentInput models.Appointment
	c.ShouldBindJSON(&appointmentInput)

	// Check if user exists
	if err := config.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User details not found"})
		return
	}

	// Create appointment instance in DB
	appointment := models.Appointment{
		Title:       appointmentInput.Title,
		Description: appointmentInput.Description,
		StartTime:   appointmentInput.StartTime,
		EndTime:     appointmentInput.EndTime,
		UserId:      int(userId),
	}

	if err := config.DB.Create(&appointment).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error creating appointment"})
		return
	}

	c.JSON(200, appointment)
}

func GetUserAppointments(c *gin.Context) {
	userId, err := middlewares.GetUserIDFromClaims(c)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "User details not found, Invalid user ID",
			"Error": err,
		})
		return
	}

	appointments := []models.Appointment{}

	// Retrieve all appointment that has user id matching userId
	if err := config.DB.Where("user_id = ?", userId).Find(&appointments).Error; err != nil {
		c.JSON(404, gin.H{"error": "User appointments not found"})
		return
	}

	c.JSON(200, appointments)
}

func GetAppointmentDetails(c *gin.Context) {
	var appointment models.Appointment

	// Retrieve appointment details
	if err := config.DB.Where("id = ?", c.Param("id")).Find(&appointment).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Appointment details not found, Invalid ID provided",
			"Error": err,
		})
		return
	}

	c.JSON(200, appointment)
}

func UpdateAppointmentDetails(c *gin.Context) {
	var existingAppointment models.Appointment
	config.DB.Where("id = ?", c.Param("id")).First(&existingAppointment)

	var input models.Appointment
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Model(&existingAppointment).Updates(input)
	c.JSON(200, existingAppointment)
}

func DeleteAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&appointment).Error; err != nil {
		c.JSON(400, gin.H{"error": "Unable to delete appointment"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Appointment deleted successfully",
	})
}
