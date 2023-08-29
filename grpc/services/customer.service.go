package services

import (
	"context"
	pd "grpc_demo/grpc/customer"
	// "grpc_demo/interfaces"
)


type RPCServer struct{
	// CustomerService interfaces.ICustomer
	pd.UnimplementedCustomerServiceServer
	
}

// func InitProfileServer(customerService interfaces.ICustomer) RPCServer{
// 	return RPCServer{customerService}
// }

func (s*RPCServer) CreateCustomer(ctx context.Context,req*pd.Customer)(*pd.CustomerResponse,error){
	return &pd.CustomerResponse{Response_Customer_ID:45674},nil
}