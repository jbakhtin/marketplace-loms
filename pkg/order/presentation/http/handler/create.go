package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
	"net/http"
)

type Item struct {
	SKU      int32  `json:"sku" validate:"required"`
	Quantity uint16 `json:"quantity" validate:"required"`
}

type CreateOrderRequest struct {
	UserID uint64 `json:"user_id" validate:"required"`
	Items  []Item `json:"items" validate:"required,dive,required"`
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
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(createOrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	orderItems := make([]models.OrderItem, len(createOrderRequest.Items))
	for i, item := range createOrderRequest.Items {
		orderItems[i].SKU = item.SKU
		orderItems[i].Count = item.Quantity
	}

	err = h.useCase.CreateOrder(r.Context(), orderItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
