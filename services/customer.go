package services

import (
	"context"
	
	
	"grpc_demo/interfaces"
	"grpc_demo/models"
	"go.mongodb.org/mongo-driver/mongo"
	
	
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx               context.Context
}

func CustomerServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{ collection , ctx}
}

func (p *CustomerService) CreateCustomer(user *models.Customer) (error) {
	 
	_, err := p.CustomerCollection.InsertOne(p.ctx, &user)

	if err != nil {
		return   err
	}
	return nil

}