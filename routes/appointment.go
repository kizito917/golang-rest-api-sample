package routes

import (
	"golang-rest-api/controllers"
	"golang-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func AppointmentRoutes(router *gin.Engine) {
	router.POST("/appointment/new", middlewares.CheckAuthorizationValidity, controllers.CreateNewAppointment)
	router.GET("/appointment/my-appointments", middlewares.CheckAuthorizationValidity, controllers.GetUserAppointments)
}
