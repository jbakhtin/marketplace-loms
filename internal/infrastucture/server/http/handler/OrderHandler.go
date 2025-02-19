package handler

import (
	"encoding/json"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/http/dto"
	"net/http"
)

type OrderHandler struct {
}

func NewOrderHandler() (OrderHandler, error) {
	return OrderHandler{}, nil
}

func (o *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var createOrderRequest dto.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&createOrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	createOrderResponse := dto.CreateOrderResponse{
		OrderID: 1, // TODO: remove constant
	}

	var buf []byte
	err = json.Unmarshal(buf, &createOrderResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(buf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (o *OrderHandler) Info(responseWriter http.ResponseWriter, request *http.Request) {

}

func (o *OrderHandler) Pay(responseWriter http.ResponseWriter, request *http.Request) {

}

func (o *OrderHandler) Cancel(responseWriter http.ResponseWriter, request *http.Request) {

}
