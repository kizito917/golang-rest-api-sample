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

func GetAppointmentDetails() {

}
