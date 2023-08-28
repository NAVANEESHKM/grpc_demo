package main

import (
	"context"
	"fmt"
	"log"
	"grpc_demo/config"
	"grpc_demo/constants"
	"grpc_demo/controllers"
	"grpc_demo/routes"
	"grpc_demo/services"

	//	"rest-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	mongoclient *mongo.Client
	ctx         context.Context
	server         *gin.Engine
)


func initApp(mongoClient *mongo.Client){
	//Customer Collection
	ctx = context.TODO()
	CustomerCollection := mongoClient.Database(constants.DatabaseName).Collection("grpc_customer")
	CustomerService := services.CustomerServiceInit(CustomerCollection, ctx)
	CustomerController := controllers.InitCustomerController(CustomerService)
	routes.CustomerRoute(server,CustomerController)

}

func main(){
	server = gin.Default()
	mongoclient,err :=config.ConnectDataBase()
	defer   mongoclient.Disconnect(ctx)
	if err!=nil{
		panic(err)
	}
	
	initApp(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))
}