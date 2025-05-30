package order

import (
	"encoding/json"
	"net/http"
)

type InfoRequest struct {
	OrderID int64
}

type InfoResponse struct {
	Status string
	User   int64
	Items  []Item
}

func (o *Handler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	o.log.Info("test", nil)

	var request InfoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	response := InfoResponse{
		Status: "test", // TODO: remove constant
		User:   1,
		Items:  make([]Item, 0),
	}

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
