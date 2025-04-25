package order

import (
	"encoding/json"
	"net/http"
)

type CreateOrderRequest struct {
	UserID uint64
	Items  []Item
}

type CreateOrderResponse struct {
	OrderID int64
}

func (o *Handler) Create(w http.ResponseWriter, r *http.Request) {
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
