package main 
import "net/http"

type handler struct {
	// gateway is going to be a layer in service discovery


}

func NewHandler () *handler {
	return &handler{}
}


func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST/api/customers/{customerID}/orders",h.)


}

func (h *handler) HandleCreateOrder (w http.ResponseWriter, r *http.Request) {

}