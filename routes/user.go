package routes

import (
	"golang-rest-api/controllers"
	"golang-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/all", middlewares.CheckAuthorizationValidity, controllers.GetUsers)
	router.GET("/:id", middlewares.CheckAuthorizationValidity, controllers.GetUserDetails)
	router.POST("/new", controllers.CreateNewUser)
	router.POST("/login", controllers.SignInUser)
	router.PUT("/:id/update", middlewares.CheckAuthorizationValidity, controllers.UpdateUser)
	router.PUT("/:id/delete", middlewares.CheckAuthorizationValidity, controllers.DeleteUser)
}
