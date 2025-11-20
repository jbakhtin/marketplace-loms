package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type InfoResponse struct {
	Status string
	User   int64
	Items  []Item
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orderIDStr := chi.URLParam(r, "OrderID")
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid order id"})
		return
	}

	_ = orderID // TODO: use orderID in logic

	// TODO: add logic
	// ...

	response := InfoResponse{
		Status: "test", // TODO: remove constant
		User:   1,
		Items:  make([]Item, 0),
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
