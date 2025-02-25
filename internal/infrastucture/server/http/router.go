package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/http/handler"
)

func NewRouter() (*chi.Mux, error) {
	orderHandler, err := handler.NewOrderHandler()
	if err != nil {
		return nil, err
	}

	stockHandler, err := handler.NewStockHandler()
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)

	router.Route("/orders", func(r chi.Router) {
		r.Post("/create", orderHandler.Create)

		r.Route("/{OrderID}", func(r chi.Router) {
			r.Get("/info", orderHandler.Info)
			r.Put("/pay", orderHandler.Pay)
			r.Put("/cancel", orderHandler.Cancel)
		})
	})

	router.Route("/stocks", func(r chi.Router) {
		r.Route("/{StockID}", func(r chi.Router) {
			r.Get("/info", stockHandler.Info)
		})
	})

	return router, nil
}
