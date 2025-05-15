package controller

import (
	"golang-rest-api/config"
	"golang-rest-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	err := config.DB.Find(&users)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to retrieve users"})
	}

	c.JSON(200, &users)
}

func CreateNewUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error creating users"})
		return
	}

	c.JSON(200, &user)
}

func GetUserDetails(c *gin.Context) {
	var user models.User
	err := config.DB.Where("id = ?", c.Param("id")).Find(&user)
	if err != nil {
		c.JSON(404, gin.H{"error": "User details not found"})
	}

	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var existingUser models.User
	config.DB.Where("id = ?", c.Param("id")).First(&existingUser)

	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Model(&existingUser).Updates(input)
	c.JSON(200, existingUser)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}
