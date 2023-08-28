package routes

import (
	"grpc_demo/controllers"

	"github.com/gin-gonic/gin"
	
)

func CustomerRoute(router *gin.Engine, controller controllers.CustomerController) {
	router.POST("/api/customer/create", controller.CreateCustomer)
	
}