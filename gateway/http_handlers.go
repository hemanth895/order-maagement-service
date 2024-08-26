package main 

import (
	"net/http"
	"errors"

	pb "github.com/hemanth895/commons/api"
	common "github.com/hemanth895/commons"
	//"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"


)

type handler struct{

	//gateway
	client 	pb.OrderServiceClient
 
}

func NewHandler(client 	pb.OrderServiceClient) *handler{
	return &handler{client}
}


func (h *handler) registerRoutes(mux *http.ServeMux){
	mux.HandleFunc("POST /api/customers/{customerID}/orders",h.HandleCreateOrder)
}


func (h *handler) HandleCreateOrder(w http.ResponseWriter ,r  *http.Request) {

	customerID := r.PathValue("customerID")
	var items []*pb.ItemsWithQuantity

	if err := common.ReadJSON(r,&items); err != nil{
		common.WriteError(w,http.StatusBadRequest,err.Error())
		return 
	}


	if err := validateItems(items);err  != nil{
		common.WriteError(w,http.StatusBadRequest,err.Error())
		return
	}



	o,err := h.client.CreateOrder(r.Context(),&pb.CreateOrderRequest{
		CustomerID:customerID,
		Items:items,
	})

	rStatus := status.Convert(err)

	if rStatus != nil{

		if rStatus.Code != nil{
			common.WriteError(w,http.StatusInternalServerError,err.Error())

			return
		}
		common.WriteError(w,http.StatusBadRequest,rStatus.Message())
		return 
	}

	if err != nil{
		
	}
	common.WriteJSON(w,http.StatusOK,o)
	
}


func validateItems(items []*pb.ItemsWithQuantity)error{
	if len(items) == 0{
		return common.ErrNoItems
		return errors.New("items must have atleast one item")
	}

	for _,i := range items{
		if i.ID == ""{
			return  errors.New("item ID is Required")
		}

		if i.Quantity <= 0 {

			return errors.New("items must ahve a valid quantity")
		}
	}

	return nil
}
