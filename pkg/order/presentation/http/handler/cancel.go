package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CancelResponse struct{}

func (h *Handler) Cancel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orderIDStr := chi.URLParam(r, "OrderID")
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_ = orderID // TODO: use orderID in logic

	// TODO: add logic
	// ...

	response := CancelResponse{}

	buf, err := json.Marshal(response)
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
