package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type InfoResponse struct {
	Count uint64
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Читаем параметр sku из query string
	skuStr := r.URL.Query().Get("sku")
	if skuStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "sku parameter is required"})
		return
	}

	sku, err := strconv.ParseInt(skuStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid sku parameter"})
		return
	}

	stockItem, err := h.useCase.StockInfo(r.Context(), int32(sku))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "stock not found"})
		return
	}

	response := InfoResponse{
		Count: stockItem.Available,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
