package interfaces

import ("grpc_demo/models"


)

type ICustomer interface{
	CreateCustomer(customer *models.Customer)(error)
}