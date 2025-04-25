package handler

import (
	"net/http"
)

type StockHandler struct {
}

func NewStockHandler() (StockHandler, error) {
	return StockHandler{}, nil
}

func (o *StockHandler) Info(writer http.ResponseWriter, response *http.Request) {

}
