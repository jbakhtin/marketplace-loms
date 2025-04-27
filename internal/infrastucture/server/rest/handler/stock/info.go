package stock

import (
	"encoding/json"
	"net/http"
)

type InfoRequest struct {
	SKU int32
}

type InfoResponse struct {
	Count int64
}

func (o *Handler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "use_case/json")

	var request InfoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	response := InfoResponse{
		Count: 1, // TODO: remove constant
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
