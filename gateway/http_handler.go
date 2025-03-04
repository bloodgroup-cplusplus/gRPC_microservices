package main

import (
	"net/http"

	pb "github.com/bloodgroup-cplusplus/gRPC_microservices/common/api"
)


type handler struct {
	// gateway is going to be a layer in service discovery
	client pb.OrderServiceClient


}


func NewHandler (client pb.OrderServiceClient) *handler {
	return &handler{client}
}


func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST/api/customers/{customerID}/orders",h.HandleCreateOrder)


}

func (h *handler) HandleCreateOrder (w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")
	

	h.client.CreateOrder(r.Context(),&pb.CreateOrderRequest{
		CustomerID:customerID ,
		Items: items,
	})
	



}