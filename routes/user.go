package routes

import (
	"golang-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/all", controller.GetUsers)
	router.GET("/:id", controller.GetUserDetails)
	router.POST("/new", controller.CreateNewUser)
	router.POST("/login", controller.SignInUser)
	router.PUT("/:id/update", controller.UpdateUser)
	router.PUT("/:id/delete", controller.DeleteUser)
}
