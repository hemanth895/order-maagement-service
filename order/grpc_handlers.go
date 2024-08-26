package  main


import (
	"context"
	"log"
	//"fmt"
	pb "github.com/hemanth895/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct{
	pb.UnimplementedOrderServiceServer

	service OrderService
}

func NewGRPCHandler(grpcServer *grpc.Server,service OrderService){

	handler := &grpcHandler{service:service}
	pb.RegisterOrderServiceServer(grpcServer,handler)
}


func (h *grpcHandler) CreateOrder(ctx context.Context,p *pb.CreateOrderRequest)(*pb.Order,error){
	


	//return nil,fmt.Errorf("some errors !!!")
	log.Printf("New order recieved! Order %v",p)

	o  := &pb.Order{
		ID:"42",
	}
	return o,nil
}


