package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func verifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}

func CheckAuthorizationValidity(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		fmt.Println("Unauthorized: No Authorization header")
		ctx.String(401, "Unauthorized")
		ctx.Abort()
		return
	}

	tokenString := authHeader
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		fmt.Println("Unauthorized")
		ctx.String(401, "Unauthorized")
		ctx.Abort()
		return
	}

	ctx.Set("claims", token.Claims)
	fmt.Printf("User authorized to execute route: %+v\n", token.Claims)

	ctx.Next()
}
