package server

import (
	"context"
	"fmt"
	"grpc_demo/constants"
	pro "grpc_demo/grpc/customer"
	rpc "grpc_demo/grpc/services"
	
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)



var (
	mongoclient *mongo.Client
	ctx         context.Context
	
)

func main(){
   lis,err:=net.Listen("tcp",constants.Port)
   if err!=nil{
	fmt.Printf("Faild to listen:%v",err)
   }
   s:=grpc.NewServer()
   pro.RegisterCustomerServiceServer(s,&rpc.RPCServer{})

   fmt.Println("Server listening")
   if err:=s.Serve(lis); err!=nil{
	fmt.Println("Server failed")
}}