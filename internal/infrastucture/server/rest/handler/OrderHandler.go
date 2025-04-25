package handler

import (
	"encoding/json"
	"net/http"
)

type Config interface {
}

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	Fatal(string, ...any)
}

type OrderHandler struct {
	cfg Config
	log Logger
}

type OrderItem struct {
	SKU   int32
	count uint16
}

type CreateOrderRequest struct {
	UserID uint64
	Items  []OrderItem
}

type CreateOrderResponse struct {
	OrderID int64
}

func NewOrderHandler(cfg Config, lgr Logger) (OrderHandler, error) {
	return OrderHandler{
		cfg: cfg,
		log: lgr,
	}, nil
}

func (o *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var createOrderRequest CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&createOrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	o.log.Info("test")

	createOrderResponse := CreateOrderResponse{
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
	o.log.Info("test", "test")
	o.log.Debug("test1", "test")
	o.log.Debug("test2", "test")
}

func (o *OrderHandler) Pay(responseWriter http.ResponseWriter, request *http.Request) {

}

func (o *OrderHandler) Cancel(responseWriter http.ResponseWriter, request *http.Request) {

}
