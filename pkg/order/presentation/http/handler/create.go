package handler

import (
	"encoding/json"
	"net/http"
)

type Item struct {
}

type CreateOrderRequest struct {
	UserID uint64
	Items  []Item
}

type CreateOrderResponse struct {
	OrderID int64
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var createOrderRequest CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&createOrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	createOrderResponse := CreateOrderResponse{
		OrderID: 1, // TODO: remove constant
	}

	buf, err := json.Marshal(createOrderResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(buf)
	if err != nil {
		return
	}
}
