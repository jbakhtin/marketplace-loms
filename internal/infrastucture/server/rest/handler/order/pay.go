package order

import (
	"encoding/json"
	"net/http"
)

type PayRequest struct {
	OrderID int64
}

type PayResponse struct{}

func (o *Handler) Pay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request PayRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	response := PayResponse{}

	var buf []byte
	err = json.Unmarshal(buf, &response)
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
