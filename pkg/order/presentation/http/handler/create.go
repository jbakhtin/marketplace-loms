package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-loms/pkg/order/app"
)

type Item struct {
	SKU      int32  `json:"sku" validate:"required,gt=0"`
	Quantity uint16 `json:"quantity" validate:"required,gt=0"`
}

type CreateOrderRequest struct {
	UserID uint64 `json:"user_id" validate:"required,gt=0"`
	Items  []Item `json:"items" validate:"required,min=1,dive"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"order_id"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var createOrderRequest CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&createOrderRequest); err != nil {
		h.logger.Error("failed to decode request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(createOrderRequest); err != nil {
		h.logger.Warn("validation failed", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Преобразуем presentation DTO в application DTO
	itemsDTO := make([]app.CreateOrderItemDTO, len(createOrderRequest.Items))
	for i, item := range createOrderRequest.Items {
		itemsDTO[i] = app.CreateOrderItemDTO{
			SKU:      item.SKU,
			Quantity: item.Quantity,
		}
	}

	orderID, err := h.useCase.CreateOrder(r.Context(), createOrderRequest.UserID, itemsDTO)
	if err != nil {
		h.logger.Error("failed to create order", "error", err, "user_id", createOrderRequest.UserID)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := CreateOrderResponse{OrderID: orderID}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("failed to encode response", "error", err)
	}
}
