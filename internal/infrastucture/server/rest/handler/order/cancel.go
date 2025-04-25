package order

import (
	"encoding/json"
	"net/http"
)

type CancelRequest struct {
	OrderID int64
}

type CancelResponse struct{}

func (o *Handler) Cancel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request CancelRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	response := CancelResponse{}

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
