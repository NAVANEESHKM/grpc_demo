package controllers

import (
	"net/http"
	"strings"

	"grpc_demo/interfaces"
	"grpc_demo/models"

	"github.com/gin-gonic/gin"
	
)

type CustomerController struct {
	CustomerService interfaces.ICustomer
}

func InitCustomerController(CustomerService interfaces.ICustomer) CustomerController {
	return CustomerController{CustomerService} //DI(dependency injection) pattern
}

func  (pc *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer *models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	err := pc.CustomerService.CreateCustomer(customer)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	
}
