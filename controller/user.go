package controller

import (
	"fmt"
	"golang-rest-api/config"
	"golang-rest-api/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = os.Getenv("JWT_SECRET")
var secretKey = []byte(jwtKey)

func createToken(name string) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": name,                             // Subject (user identifier)
		"iss": "golang-rest-api",                // Issuer
		"aud": "user",                           // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	// Print information about the created token
	fmt.Printf("Token claims added: %+v\n", claims)
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
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

	response := models.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(200, response)
}

func SignInUser(c *gin.Context) {
	var loginPayload models.UserLoginPayload
	c.BindJSON(&loginPayload)

	// Check if user email exists
	var dbUser models.User
	if err := config.DB.Where("email = ?", loginPayload.Email).First(&dbUser).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginPayload.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := createToken((dbUser.Name))
	if err != nil {
		c.JSON(400, gin.H{"error": "Unable to process login..."})
	}

	response := models.UserResponse{
		Id:    dbUser.Id,
		Name:  dbUser.Name,
		Email: dbUser.Email,
		Token: token,
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"user":    response,
	})
}

func GetUsers(c *gin.Context) {
	users := []models.User{}
	err := config.DB.Find(&users)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to retrieve users"})
	}

	c.JSON(200, &users)
}

func GetUserDetails(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).Find(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User details not found"})
		return
	}

	response := models.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(200, response)
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
