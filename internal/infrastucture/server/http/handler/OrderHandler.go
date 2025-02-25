package handler

import (
	"encoding/json"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/http/dto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

type OrderHandler struct {
	log zap.Logger
}

func NewOrderHandler() (OrderHandler, error) {
	return OrderHandler{
		log: *zap.New(zapcore.NewTee()),
	}, nil
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

	o.log.Info("test", zap.String("test", "Hello, FluentD!"))

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
	o.log.Info("test", zap.String("test", "Hello, FluentD!"))
	o.log.Debug("test1", zap.String("test2", "Hello, Debug!"))
	o.log.Debug("test2", zap.Any("test3", "Hello, Debug!"))
}

func (o *OrderHandler) Pay(responseWriter http.ResponseWriter, request *http.Request) {

}

func (o *OrderHandler) Cancel(responseWriter http.ResponseWriter, request *http.Request) {

}
