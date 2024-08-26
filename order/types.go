package  main

import (
	"context"
	pb "github.com/hemanth895/commons/api"
)

type OrderService interface {

	CreateOrder(context.Context)error
	ValidateOrder(context.Context,*pb.CreateOrderRequest)error
}

type OrdersStore interface {
	Create(context.Context)error
} 