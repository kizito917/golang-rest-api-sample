package routes

import (
	"golang-rest-api/controllers"
	"golang-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func AppointmentRoutes(router *gin.Engine) {
	router.POST("/appointment/new", middlewares.CheckAuthorizationValidity, controllers.CreateNewAppointment)
	router.GET("/appointment/my-appointments", middlewares.CheckAuthorizationValidity, controllers.GetUserAppointments)
	router.GET("/appointment/:id", middlewares.CheckAuthorizationValidity, controllers.GetAppointmentDetails)
	router.PUT("/appointment/:id/update", middlewares.CheckAuthorizationValidity, controllers.UpdateAppointmentDetails)
	router.DELETE("/appointment/:id/delete", middlewares.CheckAuthorizationValidity, controllers.DeleteAppointment)
}
